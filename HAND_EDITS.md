# Hand-Edits to Generated SDK Code

This file catalogs every commit since each version's initial generation that modified files inside `<VERSION>/api/keyfactor/v{1,2}/`. These files are nominally generator output ("DO NOT EDIT" headers), but the project's `.openapi-generator-ignore` files are empty — no protection mechanism is in place. **Naive regeneration would silently drop every hand-edit listed below.**

This catalog is read by `scripts/check-hand-edits.sh` during `make regen-diff` to verify nothing is silently lost.

## Scope

- **In scope for the upcoming v25 regen**: `v25/api/keyfactor/v1/` and `v25/api/keyfactor/v2/`.
- **Out of scope**: `v24/...` and `v2/...` — documented for posterity. No v24 swagger has been supplied, so v24 is not being regenerated in this work.

## Conventions

For each file, hand-edits are listed in commit order (oldest first). Each entry notes:

- **Commit SHA + subject** — recover the full diff with `git show <sha> -- <path>`.
- **What it changed** — brief description.
- **Reproduced by new swagger?** — Yes if the upstream swagger at `swagger/Keyfactor-Command-v25-v1.swagger.json` or `…-v2.swagger.json` contains the same shape; No if it does not.
- **Regression test pins this?** — Yes if a `*_test.go` test would fail without the hand-edit.
- **Action on regen** — `preserve` (must be re-applied post-regen), `verify` (re-check whether reproduced), `obsolete` (intentional removal), `docs-only` (no behavioral impact).

---

## v25 — IN SCOPE FOR UPCOMING REGEN

### `v25/api/keyfactor/v1/client.go` + `v25/api/keyfactor/v2/client.go`

These files contain the generator-output `APIClient` struct + transport. They have **three** distinct behavioral hand-edits across them. All are pure Go logic; **none are encoded in any swagger**. Regen will overwrite the file and all three must be re-applied.

1. **`0c1df4d`** — *fix: Use access token provided to OAuth config on validate* — adds 3 lines inside the OAuth validate path so callers supplying only `Hostname + AccessToken` (without client credentials) succeed.
   - Reproduced by swagger: **No**
   - Pinned by test: **Yes** — `229db7d` added `*_test.go` regression tests for OAuth access-token mode
   - Action: **preserve**

2. **`b374f3d`** — *Do not add port to URL if using default 443* — one-liner inside `prepareRequest()`: adds `&& serverConfig.Port != 443` to the port-append condition.
   - Reproduced by swagger: **No**
   - Pinned by test: unknown — search for tests asserting URL form
   - Action: **preserve**

3. **`229db7d`** (also adds tests) — *fix: restore AccessToken, Audience, Scopes to buildHttpClientV2 OAuth config* — re-adds 3 fields to a struct literal inside `buildHttpClientV2()` that had been accidentally stripped by `2b88eb2` (2026-03-18). Affects both v1 and v2 `client.go`.
   - Reproduced by swagger: **No**
   - Pinned by test: **Yes** — `v25/api/keyfactor/v{1,2}/client_test.go` (99 lines added in this commit)
   - Action: **preserve**

### `v25/api/keyfactor/v1/model_css_cms_core_enums_enrollment_type.go`

1. **`2d5c09e`** — *fix(enrollment type): Fix the allowed enum values for enrollment patterns* — adds `CSSCMSCOREENUMSENROLLMENTTYPE__3, __5, __6, __7` to the enum const block and to `AllowedCSSCMSCoreEnumsEnrollmentTypeEnumValues`.
   - Reproduced by new swagger: **No** — `swagger/Keyfactor-Command-v25-v1.swagger.json` defines `CSS.CMS.Core.Enums.EnrollmentType` with enum `[0, 1, 2, 4]`. Values `3, 5, 6, 7` are still missing upstream.
   - Pinned by test: not directly — but anyone using these enum values in code will break.
   - Action: **preserve** post-regen, AND file an issue/PR upstream to add the four values to the Command swagger so future regens reproduce them.

### `v25/api/keyfactor/v1/docs/CSSCMSCoreEnumsEnrollmentType.md`

1. **`2d5c09e`** — docs counterpart of the enum file above.
   - Reproduced by swagger: **No** (docs follow the swagger)
   - Action: **preserve** (regenerated from the patched enum file).

### `v25/api/keyfactor/v1/model_templates_template_retrieval_response.go`

1. **`af6340b`** *(part of `Ab#82568`)* — adds `Manageability` field (NullableInt32) with getter/setter and ToMap entry. Reason: the field was in API responses but missing from the Go struct, so it was always deserializing as zero. **Also adds CertificateCleanupEnabled, TimeAfterExpiration, TimeAfterExpirationUnits, DeleteWithArchivedKey fields**, plus nil-pointer guards on nullable getters.
   - Reproduced by new swagger: **Verify** — check whether the v25 swagger now includes `Manageability` and the cleanup fields. If yes, regen will produce them and the hand-edit is obsolete (with possible field-name diffs). If no, preserve.
   - Pinned by test: needs verification.
   - Action: **verify**, then **preserve** if not reproduced.

### `v25/api/keyfactor/v1/model_certificate_authorities_certificate_authority_request.go` and `…_response.go`

1. **`229db7d`** — adds `UseForEnrollment` and certificate-cleanup fields to both the request and response CA models. (Wait — this commit touches v24 only for these models; v25's CA models were not modified. Confirm during cataloging.)
   - Reproduced by new swagger: **No** — `swagger/Keyfactor-Command-v25-v1.swagger.json` does not include `UseForEnrollment` in `CertificateAuthoritiesCertificateAuthorityRequest.properties`.
   - **Note**: this is a v24 hand-edit. Whether v25 needs the same depends on what `229db7d` actually changed in v25.
   - Action: **verify**, then preserve in v25 too if missing upstream.

### `v25/api/keyfactor/v1/client_test.go`, `v25/api/keyfactor/v2/client_test.go`

1. **`229db7d`** — adds 99 lines of regression tests for OAuth access-token mode and CA cleanup field behavior.
   - These are **regression tests** (read-only per project rule). Regen must not overwrite or remove them.
   - Action: **preserve verbatim**. The check script must error if these are missing from `.regen-staging/`.

### `v25/api/keyfactor/v1/model_certificate_authorities_test.go`

1. **`229db7d`** — adds 264 lines of regression tests for the CA model hand-edits.
   - Same constraint as above — read-only regression tests.
   - Action: **preserve verbatim**.

### `v25/api/keyfactor/v1/README.md`, `v25/api/keyfactor/v2/README.md`

1. **`91c470e`** — minor documentation fix (3 lines per file).
   - Reproduced by swagger: **No** (generated from the swagger description fields)
   - Action: **docs-only** — re-apply if it still makes sense after regen, otherwise drop.

---

## v24 — OUT OF SCOPE FOR THIS REGEN (documented for posterity)

v24 has substantially more hand-editing than v25 because it has been in maintenance longer. Listed by file, in commit order:

### `v24/api/keyfactor/v1/client.go` (9 commits since initial)
- All three behaviors documented in the v25 section above plus historical v24-only fixes. Most importantly, `af6340b` adds the **entire `NewAPIClientWithAuth` function** (~60 lines) used for VCR test injection. This function exists in v24 only as a hand-edit.

### `v24/api/keyfactor/v1/model_templates_template_retrieval_response.go` (4 commits)
- Same Manageability + cleanup fields hand-edits as v25.

### `v24/api/keyfactor/v1/model_templates_template_update_request.go` (3 commits)
- CertificateCleanupEnabled, TimeAfterExpiration, TimeAfterExpirationUnits, DeleteWithArchivedKey fields.

### `v24/api/keyfactor/v1/model_css_cms_core_enums_enrollment_type.go` (3 commits)
- Same `3, 5, 6, 7` enum values as v25 + initial v24-only fixes.

### `v24/api/keyfactor/v1/model_css_cms_core_enums_key_retention_policy.go` (3 commits)
### `v24/api/keyfactor/v1/model_css_cms_data_model_enums_certificate_cleanup_time_units.go`
- Created entirely by hand in `af6340b` (the enum doesn't exist in the v24 swagger).

### `v24/api/keyfactor/v1/model_certificate_authorities_certificate_authority_request.go` and `…_response.go`
- `UseForEnrollment` + cleanup fields. Added in `229db7d`.

### `v24/api/keyfactor/v1/client_test.go`, `v24/api/keyfactor/v2/client_test.go`, `v24/api/keyfactor/v1/model_certificate_authorities_test.go`
- Regression tests (`229db7d`). Read-only.

---

## Pre-existing untracked work in v24 working tree

Discovered during Phase 0:

- `v24/api/keyfactor/v1/api_application.go` (490 lines)
- `v24/api/keyfactor/v1/model_applications_application_detail_response.go`
- `v24/api/keyfactor/v1/model_applications_application_list_response.go`
- `v24/api/keyfactor/v1/model_applications_application_request.go`
- `v24/go.mod`, `v24/go.sum` modifications

These are openapi-generator output for `/Applications` endpoint coverage — produced by someone who ran a v24 regen against a swagger that included Applications, but never committed. They are **uncommitted and out of scope** for the v25 regen. The same regen against `swagger/Keyfactor-Command-v25-v1.swagger.json` will produce the v25 equivalent.

If v24 is later regenerated for the same purpose, these files document what the output looked like.

---

## Resolving spec gaps (RULE)

**Any gap between the swagger file and what we know the Command API actually accepts must be resolved by consulting the official Keyfactor Command API documentation, not by preserving the hand-edit forever.**

Process for each "Reproduced by swagger: **No**" entry above where the hand-edit looks like a missing-in-spec value (not a pure client-side fix):

1. **Look up the official Keyfactor API documentation** for the relevant endpoint/model/enum. Capture the canonical field names, types, and enum values.
2. **Compare with the hand-edit and the current swagger.** If docs agree with the hand-edit → swagger is wrong, patch the swagger file. If docs disagree with both → flag for human resolution. If docs agree with swagger → hand-edit is wrong; drop it.
3. **Patch `swagger/Keyfactor-Command-v25-v{1,2}.swagger.json` in place** so the next regen produces correct output without post-regen patching. Document the patch in this file under a new "Spec patches applied" section.
4. **Only after the swagger is canonical**, run the regen flow.

Pure client-side hand-edits (port 443 logic, OAuth access-token handling, `NewAPIClientWithAuth` for VCR tests) are **not** spec gaps — they're behavioral fixes that don't belong in the swagger and must be re-applied as post-regen patches.

## Open questions

1. **EnrollmentType enum values `3, 5, 6, 7`** — current swagger has `[0, 1, 2, 4]`. Verify against official Keyfactor API docs which values are real; patch swagger accordingly.

2. **CA model `UseForEnrollment` + certificate-cleanup fields** — verify against docs; patch swagger if confirmed.

3. **Template model `Manageability` field + cleanup field family** — verify against docs; patch swagger if confirmed.

4. **`NewAPIClientWithAuth` function** — does v25 have it (perhaps in `v25/client.go` root)? Search before treating as a missing hand-edit.

All four must be resolved before `regen-apply` runs.

## Spec patches applied

*(populated as patches are made to the swagger files during this work)*

| Date | Spec file | Patch | Source doc | Hand-edit obsoleted |
|---|---|---|---|---|
| — | — | — | — | — |
