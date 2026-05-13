#!/usr/bin/env bash
#
# check-hand-edits.sh — verify hand-edits in regen output match the manifest.
#
# Used by the SDK Makefile regen workflow. Two modes:
#
#   check <v1_stage_dir> <v2_stage_dir>
#       For each manifest row:
#         - preserve:      grep pattern in staged file; ERROR if missing
#         - verify:        grep pattern in staged file; WARN  if missing
#         - preserve-file: ERROR if file missing from staging
#         - verify-file:   WARN  if file missing from staging
#       Exits non-zero if any 'preserve' or 'preserve-file' assertion fails.
#
#   preserve <v1_baseline_dir> <v2_baseline_dir> <v1_stage_dir> <v2_stage_dir>
#       Copies every 'preserve-file' row from baseline into staging so the
#       subsequent swap to live does not wipe regression tests etc.
#       Returns non-zero only on copy failure.
#
# Manifest format: scripts/hand-edits-manifest.tsv (tab-separated, see header).

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MANIFEST="${SCRIPT_DIR}/hand-edits-manifest.tsv"

if [[ ! -f "${MANIFEST}" ]]; then
    echo "ERROR: manifest not found at ${MANIFEST}" >&2
    exit 2
fi

MODE="${1:-check}"
shift || true

# ----- read manifest into arrays of valid (non-comment, non-empty) rows -----
# Each row: category<TAB>version_api<TAB>file<TAB>pattern<TAB>ref
declare -a ROWS
while IFS= read -r line; do
    # strip CR if present (Windows-edited files)
    line="${line%$'\r'}"
    # skip comments and empty lines
    [[ -z "${line}" || "${line}" =~ ^# ]] && continue
    ROWS+=("${line}")
done < "${MANIFEST}"

# Resolve a manifest row's actual staged file path.
# $1 = v1_stage_dir, $2 = v2_stage_dir, $3 = version_api (v1|v2), $4 = file
stage_path() {
    local v1d="$1" v2d="$2" ver="$3" file="$4"
    case "${ver}" in
        v1) echo "${v1d}/${file}" ;;
        v2) echo "${v2d}/${file}" ;;
        *)  echo "" ;;
    esac
}

# Resolve a manifest row's baseline file path (for preserve mode).
baseline_path() {
    local b1d="$1" b2d="$2" ver="$3" file="$4"
    case "${ver}" in
        v1) echo "${b1d}/${file}" ;;
        v2) echo "${b2d}/${file}" ;;
        *)  echo "" ;;
    esac
}

# ----- mode: check -----
mode_check() {
    local v1_stage="${1:?v1_stage_dir required}"
    local v2_stage="${2:?v2_stage_dir required}"

    local errors=0 warnings=0 oks=0
    for row in "${ROWS[@]}"; do
        IFS=$'\t' read -r category version_api file pattern ref <<< "${row}"
        local path
        path="$(stage_path "${v1_stage}" "${v2_stage}" "${version_api}" "${file}")"
        if [[ -z "${path}" ]]; then
            echo "  ! unknown version '${version_api}' in row: ${row}" >&2
            errors=$((errors+1))
            continue
        fi

        case "${category}" in
            preserve)
                if [[ ! -f "${path}" ]]; then
                    echo "  ✗ MISSING FILE (preserve)  ${path}  [${ref}]"
                    errors=$((errors+1))
                elif grep -qF -- "${pattern}" "${path}"; then
                    oks=$((oks+1))
                else
                    echo "  ✗ MISSING PATTERN (preserve)  ${path}  pattern='${pattern}'  [${ref}]"
                    errors=$((errors+1))
                fi
                ;;
            verify)
                if [[ ! -f "${path}" ]]; then
                    echo "  ⚠ missing file (verify)  ${path}  [${ref}]"
                    warnings=$((warnings+1))
                elif grep -qF -- "${pattern}" "${path}"; then
                    oks=$((oks+1))
                else
                    echo "  ⚠ missing pattern (verify)  ${path}  pattern='${pattern}'  [${ref}]"
                    warnings=$((warnings+1))
                fi
                ;;
            preserve-file)
                if [[ -f "${path}" ]]; then
                    oks=$((oks+1))
                else
                    echo "  ✗ MISSING FILE (preserve-file)  ${path}  [${ref}]"
                    errors=$((errors+1))
                fi
                ;;
            verify-file)
                if [[ -f "${path}" ]]; then
                    oks=$((oks+1))
                else
                    echo "  ⚠ missing file (verify-file)  ${path}  [${ref}]"
                    warnings=$((warnings+1))
                fi
                ;;
            *)
                echo "  ! unknown category '${category}' in row: ${row}" >&2
                errors=$((errors+1))
                ;;
        esac
    done

    echo "  summary: ${oks} ok, ${warnings} warnings, ${errors} errors"
    if (( errors > 0 )); then
        echo "  HAND-EDIT CHECK FAILED — review HAND_EDITS.md and either patch swagger, apply post-regen patch, or update manifest."
        return 1
    fi
    if (( warnings > 0 )); then
        echo "  hand-edit warnings present — review before regen-apply."
    fi
    return 0
}

# ----- mode: preserve -----
# Copies every 'preserve-file' row from baseline into staging.
mode_preserve() {
    local v1_baseline="${1:?v1_baseline_dir required}"
    local v2_baseline="${2:?v2_baseline_dir required}"
    local v1_stage="${3:?v1_stage_dir required}"
    local v2_stage="${4:?v2_stage_dir required}"

    local copies=0 missing=0
    for row in "${ROWS[@]}"; do
        IFS=$'\t' read -r category version_api file pattern ref <<< "${row}"
        [[ "${category}" == "preserve-file" ]] || continue
        local src dst
        src="$(baseline_path "${v1_baseline}" "${v2_baseline}" "${version_api}" "${file}")"
        dst="$(stage_path   "${v1_stage}"    "${v2_stage}"    "${version_api}" "${file}")"
        if [[ ! -f "${src}" ]]; then
            echo "  ! baseline missing: ${src}  [${ref}]" >&2
            missing=$((missing+1))
            continue
        fi
        mkdir -p "$(dirname "${dst}")"
        cp "${src}" "${dst}"
        copies=$((copies+1))
        echo "  + preserved: ${dst}"
    done
    echo "  summary: ${copies} files preserved, ${missing} baseline-missing"
    (( missing == 0 ))
}

case "${MODE}" in
    check)    mode_check    "$@" ;;
    preserve) mode_preserve "$@" ;;
    *)
        echo "usage: $0 check    <v1_stage_dir> <v2_stage_dir>"
        echo "       $0 preserve <v1_baseline_dir> <v2_baseline_dir> <v1_stage_dir> <v2_stage_dir>"
        exit 2
        ;;
esac
