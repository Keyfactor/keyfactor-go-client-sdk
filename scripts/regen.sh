#!/usr/bin/env bash
#
# scripts/regen.sh — regenerate the v25 SDK from committed swagger inputs.
#
# Adapted from Keyfactor/API-definitions @ 04b7b2c (go/command/openapi-generate.sh)
# with these changes:
#   - Reads from this repo's committed swagger/ files instead of fetching
#     via curl from a live Command instance.
#   - Writes into ${OUTPUT_DIR}/ (default .regen-staging/) instead of in-place
#     into the version dir, so the safety harness (regen-prepare/diff/apply
#     in the Makefile) can review changes before swap.
#   - Uses java -jar openapi-generator-cli.jar (downloaded by the Makefile)
#     instead of an externally-installed `openapi-generator` CLI.
#
# All idempotent swagger patches from the canonical script are preserved:
#   - Add `schema: {type: string}` to POST parameters missing it
#   - Replace server URL with http://keyfactor.example.com
#   - Set info.version to the API version (1, 2)
#   - Rewrite operationIds for cleaner generated method names
#   - Strip `Keyfactor.Web.KeyfactorApi.Models.` namespace from schemas
#   - Add x-enum-varnames to ClaimType, DayOfWeek
#   - Hotfix: CSS.CMS.Core.Enums.EnrollmentType enum = [0..7] (Zendesk 139784)
#
# Patches applied by this repo's spec gaps (already in committed swagger):
#   - Header parameters missing `schema` field (475 total across v1+v2)
#   - These are idempotent with the canonical script's POST-param patch.
#
# Usage:
#   COMMAND_VERSION=v25
#   COMMAND_VERSION_FULL=25.1.1
#   OUTPUT_DIR=.regen-staging
#   SWAGGER_V1=swagger/Keyfactor-Command-v25-v1.swagger.json
#   SWAGGER_V2=swagger/Keyfactor-Command-v25-v2.swagger.json
#   GENERATOR_JAR=openapi-generator-cli.jar
#   scripts/regen.sh

set -euo pipefail

# ============== Inputs (override via env) ==============
COMMAND_VERSION="${COMMAND_VERSION:-v25}"
COMMAND_VERSION_FULL="${COMMAND_VERSION_FULL:-25.1.1}"
OUTPUT_DIR="${OUTPUT_DIR:-.regen-staging}"
SWAGGER_V1="${SWAGGER_V1:-swagger/Keyfactor-Command-${COMMAND_VERSION}-v1.swagger.json}"
SWAGGER_V2="${SWAGGER_V2:-swagger/Keyfactor-Command-${COMMAND_VERSION}-v2.swagger.json}"
GENERATOR_JAR="${GENERATOR_JAR:-openapi-generator-cli.jar}"
API_VERSIONS=("1" "2")

# ============== Pre-flight ==============
for dep in jq sed java python3; do
  command -v "$dep" >/dev/null 2>&1 || { echo "ERROR: '$dep' not on PATH" >&2; exit 1; }
done
[[ -f "${GENERATOR_JAR}" ]] || { echo "ERROR: ${GENERATOR_JAR} not found (run 'make' to download)" >&2; exit 1; }
[[ -f "${SWAGGER_V1}" ]] || { echo "ERROR: ${SWAGGER_V1} not found" >&2; exit 1; }
[[ -f "${SWAGGER_V2}" ]] || { echo "ERROR: ${SWAGGER_V2} not found" >&2; exit 1; }
[[ -d custom-templates/go ]] || { echo "ERROR: custom-templates/go not found" >&2; exit 1; }
[[ -f openapi-config.yml ]] || { echo "ERROR: openapi-config.yml not found" >&2; exit 1; }

base_dir="$(pwd)"
custom_templates_path="custom-templates/go"
config_file_name="openapi-config.yml"
current_year="$(date +%Y)"

mkdir -p "${OUTPUT_DIR}/${COMMAND_VERSION}"

# Dynamically generated variables populated by API client version (for postprocess)
imports=""
client_instantiation=""
client_list=""
documentation_list=""
api_client_examples=""

# ============== Per-version generation ==============
for version_num in "${API_VERSIONS[@]}"; do
  cd "${base_dir}"
  version="v${version_num}"
  api_version_path="${OUTPUT_DIR}/${COMMAND_VERSION}/api/keyfactor/${version}"
  echo ">>> ${version}: output -> ${api_version_path}"
  mkdir -p "${api_version_path}"

  if [[ "${version_num}" == "1" ]]; then
    source_spec="${SWAGGER_V1}"
  else
    source_spec="${SWAGGER_V2}"
  fi

  spec_file="${OUTPUT_DIR}/.spec_${version}.json"
  spec_raw_file="${OUTPUT_DIR}/.spec_${version}-raw.json"

  echo "    using committed swagger ${source_spec} as raw input"
  cp "${source_spec}" "${spec_file}"
  cp "${source_spec}" "${spec_raw_file}"

  # --- patch 1: ensure POST parameters have schema (idempotent) ---
  jq '
    .paths |= with_entries(
      if .value.post and .value.post.parameters then
        .value.post.parameters |= map(
          if has("schema") then . else . + { "schema": { "type": "string" } } end
        )
      else .
      end
    )
  ' "${spec_file}" > "${spec_file}.tmp" && mv "${spec_file}.tmp" "${spec_file}"

  # --- patch 2: replace server URL ---
  jq '.servers |= map(.url = "http://keyfactor.example.com")' \
    "${spec_file}" > "${spec_file}.tmp" && mv "${spec_file}.tmp" "${spec_file}"

  # --- patch 3: set info.version to the API version ---
  jq --arg version "${version_num}" '.info.version = $version' \
    "${spec_file}" > "${spec_file}.tmp" && mv "${spec_file}.tmp" "${spec_file}"

  # --- patch 4: rewrite operationIds for cleaner method names ---
  jq --arg version "${version_num}" '
  .paths |= with_entries(
    . as $path |
    .value |= with_entries(
        .value.operationId = (
          (if .key == "post" then "Create"
           elif .key == "put" then "Update"
           elif .key == "delete" then "Delete"
           elif .key == "patch" then "Patch"
           else "Get" end) +
          (
            $path.key | gsub("[{}]"; "") | split("/") | map(select(. != "")) | map(
              if . == "id" then "ById" else . end
            ) | join("_")
          )
        )
    )
  )' "${spec_file}" > "${spec_file}.tmp" && mv "${spec_file}.tmp" "${spec_file}"

  # --- patch 5: strip Keyfactor.Web.KeyfactorApi.Models. namespace ---
  sed 's/Keyfactor\.Web\.KeyfactorApi\.Models\.//g' \
    "${spec_file}" > "${spec_file}.tmp" && mv "${spec_file}.tmp" "${spec_file}"

  # --- patch 6: x-enum-varnames on common enums ---
  jq '
    .components.schemas."CSS.CMS.Core.Enums.ClaimType" |=
     if has("x-enum-varnames") then .
     else (. + {"x-enum-varnames": ["User", "Group", "Computer", "OAuthOid", "OAuthRole", "OAuthSubject", "OAuthClientId" ] })
     end
  ' "${spec_file}" > "${spec_file}.tmp" && mv "${spec_file}.tmp" "${spec_file}"

  jq '
    .components.schemas."System.DayOfWeek" |=
     if has("x-enum-varnames") then .
     else (. + {"x-enum-varnames": ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday" ] })
     end
  ' "${spec_file}" > "${spec_file}.tmp" && mv "${spec_file}.tmp" "${spec_file}"

  # --- patch 7: EnrollmentType enum hotfix (Zendesk 139784) ---
  jq '
    if .components.schemas["CSS.CMS.Core.Enums.EnrollmentType"]
    then .components.schemas["CSS.CMS.Core.Enums.EnrollmentType"].enum = [0, 1, 2, 3, 4, 5, 6, 7]
    else . end
  ' "${spec_file}" > "${spec_file}.tmp" && mv "${spec_file}.tmp" "${spec_file}"

  echo "    patched swagger written to ${spec_file}"

  # ============== Run generator ==============
  openapi_generator_version="$(java -jar "${GENERATOR_JAR}" version 2>/dev/null | tail -1)"

  echo "    running openapi-generator into ${api_version_path}"
  # NOTE: nameMappings (HasValue -> DoesHaveValue) is provided by
  # openapi-config.yml; it is honored by openapi-generator 7.x but silently
  # ignored by 6.x. The Makefile pin (OPENAPI_GENERATOR_VERSION) must be 7.x
  # or later, otherwise the generated v1 fails to compile due to a
  # HasValue field/method collision in KeyfactorSecret.
  java -jar "${GENERATOR_JAR}" generate \
    -g go \
    -i "${spec_file}" \
    -t "${custom_templates_path}" \
    -c "${config_file_name}" \
    -o "${api_version_path}" \
    -p sdkVersion="${COMMAND_VERSION}" \
    -p apiVersion="${COMMAND_VERSION_FULL}" \
    -p apiNameSuffix="Api" \
    -p packageName="${version}" \
    -p openApiVersion="${openapi_generator_version}" \
    -p currentYear="${current_year}" \
    -p isGoSubmodule=false \
    -p disallowAdditionalPropertiesIfNotPresent=false \
    --git-user-id "Keyfactor" \
    --git-repo-id "keyfactor-go-client-sdk" \
    >/dev/null

  echo "    generator complete"

  # Post-generator cleanup inside the version directory
  cd "${api_version_path}"
  rm -rf ./test || true
  rm -f .travis.yml git_push.sh .openapi-generator-ignore || true

  # Copy config.yml into the version directory (matches canonical layout)
  cp "${base_dir}/config.yml" .

  # gofmt + goimports if available
  if command -v goimports >/dev/null 2>&1; then
    goimports -w . || true
  fi
  if command -v gofmt >/dev/null 2>&1; then
    gofmt -w . || true
  fi

  # The canonical script removes go.mod/go.sum here too — generator emits per-package
  # ones we don't want. The version root carries the real go.mod/go.sum.
  rm -f go.mod go.sum || true

  cd "${base_dir}"

  # Accumulate postprocess substitution variables
  imports+="${version} \"github.com/Keyfactor/keyfactor-go-client-sdk/${COMMAND_VERSION}/api/keyfactor/${version}\"\n"
  client_instantiation+="clientV${version_num}, err := ${version}.NewAPIClient(cfg)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\n\t"
  client_list+="V${version_num}: clientV${version_num},\n\t\t"
  documentation_list+="- [${version} API Documentation](./api/keyfactor/${version}/README.md)\n"
  api_client_examples+="// api := client.V${version_num}.ExampleApi // Access V${version_num} API Service\n// req := api.NewUpdateExampleRequest(ctx) // Build new API request\n// req = req.ExampleUpdateRequest(${version}.ExampleUpdateRequest{ Name: \"Hello\" }) // Add a body to the API request (if applicable)\n// resp, httpResp, err := req.Execute() // Execute request\n\n"
done

# ============== Post-process: handwritten wrappers at version root ==============
echo ">>> Post-processing: writing handwritten wrappers at ${OUTPUT_DIR}/${COMMAND_VERSION}/"
cd "${OUTPUT_DIR}/${COMMAND_VERSION}"

cp "${base_dir}/custom-templates/postprocess/go/go.mod.template" go.mod
cp "${base_dir}/custom-templates/postprocess/go/go.sum.template" go.sum
cp "${base_dir}/custom-templates/postprocess/go/client.go.template" client.go
cp "${base_dir}/custom-templates/postprocess/go/README.md.template" README.md
cp "${base_dir}/custom-templates/postprocess/go/helpers.go.template" helpers.go

# Substitute template variables. Use python because sed with multi-line + complex
# strings is fragile across BSD/GNU. Python keeps newlines and special chars intact.
python3 - "$COMMAND_VERSION" "$COMMAND_VERSION_FULL" "$imports" "$client_instantiation" "$client_list" "$documentation_list" "$api_client_examples" <<'PY'
import sys, pathlib, codecs

ver, ver_full, imports, client_inst, client_list, doc_list, api_examples = sys.argv[1:]
# Unescape \n / \t embedded in the bash-built strings
def unesc(s):
    return codecs.decode(s, 'unicode_escape')

mappings = {
    "{{packageVersion}}": ver,
    "{{version}}": ver_full,
    "{{imports}}": unesc(imports).rstrip(),
    "{{clientInstantiation}}": unesc(client_inst).rstrip(),
    "{{clientList}}": unesc(client_list).rstrip(),
    "{{documentationList}}": unesc(doc_list).rstrip(),
    "{{apiClientExamples}}": unesc(api_examples).rstrip(),
}

for path in ["go.mod", "go.sum", "client.go", "helpers.go", "README.md"]:
    p = pathlib.Path(path)
    if not p.exists():
        continue
    text = p.read_text()
    for k, v in mappings.items():
        text = text.replace(k, v)
    p.write_text(text)
    print(f"    substituted: {path}")
PY

if command -v go >/dev/null 2>&1; then
  go mod tidy >/dev/null 2>&1 || echo "    WARN: 'go mod tidy' failed (likely safe in staging)"
  go fmt . >/dev/null 2>&1 || true
fi

cd "${base_dir}"

echo ">>> Regen complete. Output: ${OUTPUT_DIR}/${COMMAND_VERSION}/"
