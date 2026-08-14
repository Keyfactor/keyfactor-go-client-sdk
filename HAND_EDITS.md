# Hand-Edits to Generated SDK Code

## Scope

This file catalogs commits that modified files inside **`v24/api/keyfactor/v{1,2}/`** after
their initial generation (baseline: `2ed41db`, "Generate V24 client"). These files carry
generator-output "DO NOT EDIT" headers, but the project's `.openapi-generator-ignore` files
are 1040 bytes of generator boilerplate comments with no active ignore patterns (verified by
byte count and content inspection) — no protection mechanism is in place. **Without the right
templates and swagger patches, naive regeneration of v24 would silently drop every hand-edit
listed below.**

**This document does NOT currently cover `v25/api/keyfactor/v{1,2}/`, `v2/api/keyfactor/`, or
the repo-root `api/keyfactor/`.** `v25` has its own hand-edit history starting from its own
generation baseline (`536f3e2`, "feat(api): Add support for Keyfactor API Command up to
25.1.1") that has not been fully audited against this convention — see the "v25 (partial)"
section below for the one edit ported there as part of this change set. Treat any claim of
completeness in this file as scoped to v24 only; do not assume v25/v2/root are safe to
regenerate just because they aren't listed here.

## Conventions

For each file, hand-edits are listed in commit order (oldest first). Each entry notes:

- **Commit SHA + subject** — recover the full diff with `git show <sha> -- <path>`.
- **What it changed** — brief description.
- **Reproduced by upstream swagger?** — this repo has no swagger/OpenAPI spec file checked in
  at any commit in its history (`find . -iname '*swagger*' -o -iname '*openapi*'` under
  `v24/` returns nothing but generated output and `.openapi-generator-ignore`), so the answer
  is **No** for every entry below: none of these edits can be verified against a source spec
  we can regenerate from. "No" here does not mean "wrong," only "pure Go logic (or a
  hand-authored file matching generator conventions) with no committed spec counterpart."
- **Regression test pins this?** — Yes if a `*_test.go` test would fail (or fail to compile)
  without the hand-edit; "compile-only" if removal would only break a build elsewhere in this
  repo without a test asserting behavior.
- **Action on regen** — `preserve` (must be re-applied post-regen), `verify` (re-check whether
  reproduced), `obsolete` (intentional removal), `docs-only` (no behavioral impact).

---

## v24

### `v24/api/keyfactor/v1/client.go` + `v24/api/keyfactor/v2/client.go`

1. **`af6340b`** — *Ab#82568 (#30)* — adds `NewAPIClientWithAuth(auth AuthConfig) *APIClient`
   to both files: constructs an `APIClient` with a pre-built `AuthConfig`, bypassing the
   network call inside `Authenticate()`. This is the entry point the VCR-cassette-based unit
   test harness (in this repo and in consumers such as the Terraform provider) uses to inject
   a fake or replay-mode `AuthConfig` without hitting a real Command server.
   - Reproduced by upstream swagger: **No**.
   - Regression test pins this: **Compile-only** — `v24/client.go` (the top-level wrapper)
     calls `v1.NewAPIClientWithAuth` / `v2.NewAPIClientWithAuth` directly, so `go build ./...`
     fails if either is removed or its signature changes, but no test in *this* repo asserts
     its runtime behavior. Downstream consumers' VCR test suites depend on it at compile time
     too.
   - Action: **preserve**.

2. **`229db7d`** — *Feat/ca cleanup enrollment fields (#32)* — restores `AccessToken`,
   `Audience`, and `Scopes` to the `auth_providers.CommandConfigOauth{...}` struct literal
   inside `buildHttpClientV2`'s OAuth branch. Commit `2b88eb2` (2026-03-18) had silently
   dropped these three fields in an earlier refactor, which broke pre-fetched
   `access_token`-only authentication (callers supplying just hostname + access token, no
   `client_id`/`client_secret`/`token_url`).
   - Reproduced by upstream swagger: **No**.
   - Regression test pins this: **Yes** —
     `TestCommandConfigOauth_AccessTokenFieldPropagation` fails to compile if any of the three
     fields are removed from either struct, and asserts they propagate correctly.
   - Action: **preserve**. (This is the second time these fields were silently dropped and
     restored — see the near-identical prior loss at `2b88eb2` — so it is a high-value
     preserve.)

3. **`229db7d`** (same commit) — changes `prepareRequest`'s port guard from
   `serverConfig.Port > 0 && serverConfig.Port <= 65535` to
   `serverConfig.Port > 0 && serverConfig.Port <= 65535 && serverConfig.Port != 443`, so a
   `Server` configured with the default HTTPS port no longer produces
   `https://host:443/...` request URLs (some servers/proxies that match on an exact,
   port-suffix-free `Host` header rejected the explicit `:443`).
   - Reproduced by upstream swagger: **No**.
   - Regression test pins this: **Yes, as of this change set** —
     `TestPrepareRequest_Port443Guard` (added alongside this document update) asserts both
     that port 443 is omitted and that a non-443 port (8443) is still appended. Prior to this
     change, reverting the guard failed nothing in CI.
   - Action: **preserve**.

4. **`2a6c5b4`** — *fix(v24): plumb Server.ClientTimeout into rebuilt auth config* — inside
   `buildHttpClientV2()`, adds `HttpClientTimeout: cfg.ClientTimeout` to the
   `baseConfig := auth_providers.CommandAuthConfig{...}` struct literal in both files. Without
   it, `Server.ClientTimeout` (added upstream by `keyfactor-auth-client-go` v1.6.0-rc.2 to fix
   [issue #51](https://github.com/Keyfactor/keyfactor-auth-client-go/issues/51)) was silently
   dropped when this SDK rebuilt its own `CommandAuthConfig`, so every caller — including the
   Terraform provider's `request_timeout` setting — fell back to
   `auth_providers.DefaultClientTimeout` (60s) regardless of what was configured. This
   surfaced as `net/http: timeout awaiting response headers` on long-running calls such as PFX
   enrollment.
   - Reproduced by upstream swagger: **No**.
   - Regression test pins this: **Yes** — `TestBuildHttpClientV2_ClientTimeoutPropagation` in
     both `v1/client_test.go` and `v2/client_test.go` calls `buildHttpClientV2()` against a
     fake Command server and asserts the resulting `CommandAuthConfigBasic.HttpClientTimeout`
     and derived `BuildTransport().ResponseHeaderTimeout` reflect the configured value. That
     test is itself hermetic against ambient `KEYFACTOR_SKIP_VERIFY` /
     `KEYFACTOR_CA_CERT` / `KEYFACTOR_CLIENT_TIMEOUT` env values (see the `unsetEnvForTest`
     helper in the same file) so it can't be spuriously broken by a caller's shell
     environment.
   - Action: **preserve**.

### `v24/api/keyfactor/v1/model_certificate_authorities_certificate_authority_request.go` + `..._response.go`

5. **`229db7d`** — adds `UseForEnrollment *bool`, `CertificateCleanupEnabled NullableBool`,
   `DeleteWithArchivedKey NullableBool`, `TimeAfterExpiration NullableInt32`, and
   `TimeAfterExpirationUnits *CSSCMSDataModelEnumsCertificateCleanupTimeUnits` fields (with
   generator-style getters/setters/Has*) to both the CA request and response models.
   - Reproduced by upstream swagger: **No**.
   - Regression test pins this: **Yes** — `TestCARequestFields_CleanupAndEnrollment`,
     `TestCAResponseFields_CleanupAndEnrollment`, `TestCARequestFields_NilSafety`, and
     `TestCAResponseFields_NilSafety` in `model_certificate_authorities_test.go` exercise every
     setter/getter/Has* and nil-receiver safety for all five fields on both structs.
   - Action: **preserve**.

### `v24/api/keyfactor/v1/model_css_cms_core_enums_enrollment_type.go`

6. **`af6340b`** — adds enum value `3` and the bitmask-combination values `5`, `6`, `7` to
   `AllowedCSSCMSCoreEnumsEnrollmentTypeEnumValues` (the type is a bitmask: 1=PFX, 2=CSR, so
   combined values are valid), with a comment documenting the bitmask semantics.
   - Reproduced by upstream swagger: **No**.
   - Regression test pins this: **No** — no test in this repo references
     `CSSCMSCoreEnumsEnrollmentType`'s allowed-values list.
   - Action: **preserve**.

### `v24/api/keyfactor/v1/model_css_cms_core_enums_key_retention_policy.go`

7. **`af6340b`** — rewrites `UnmarshalJSON` to try integer form first (original generated
   behavior), then fall back to a new `keyRetentionPolicyStringToInt` map for EJBCA's
   string-form responses (e.g. `"None"`, `"ShortTerm"`), defaulting unknown strings to `0`
   (`None`) rather than erroring, to avoid breaking reads of new/unknown values.
   - Reproduced by upstream swagger: **No**.
   - Regression test pins this: **No** — no test in this repo exercises
     `CSSCMSCoreEnumsKeyRetentionPolicy.UnmarshalJSON`'s string-form or unknown-value paths.
   - Action: **preserve**.

### `v24/api/keyfactor/v1/model_css_cms_data_model_enums_certificate_cleanup_time_units.go`

8. **`af6340b`** — new file (does not exist before this commit). Hand-authored in the exact
   style of a generated enum file (including the generator's `DO NOT EDIT` header and the
   known openapi-generator template artifact where `Parse()` unconditionally returns an error
   for enums generated with no string-value mapping, before checking a `stringsToEnum` map
   that is always empty — this is not new breakage, it reproduces a real upstream generator
   quirk found on comparable enum files elsewhere in this SDK). Defines
   `CSSCMSDataModelEnumsCertificateCleanupTimeUnits` (0=Days, 1=Weeks, 2=Months), consumed by
   the CA and template cleanup fields added in the same commit / in `229db7d`.
   - Reproduced by upstream swagger: **No** — hand-created, not generated.
   - Regression test pins this: **Compile-only** — `TestCARequestFields_CleanupAndEnrollment`
     references the `CSSCMSDATAMODELENUMSCERTIFICATECLEANUPTIMEUNITS__1` constant, so removing
     the type breaks compilation, but no test exercises its `UnmarshalJSON`/`Parse` behavior.
   - Action: **preserve**.

### `v24/api/keyfactor/v1/model_system_day_of_week.go` + `v24/api/keyfactor/v2/model_system_day_of_week.go`

9. **`96dd817`** — rewrites `SystemDayOfWeek.UnmarshalJSON` to try the integer form first
   (preserving the original generated validation against
   `AllowedSystemDayOfWeekEnumValues`), then fall back to the existing `Parse()` day-name
   mapping (e.g. `"Monday"`) when the payload is a JSON string, since Keyfactor Command
   serializes `WeeklyModel.Days` as day-name strings in some API responses. Applied
   identically to both v1 and v2 packages.
   - Reproduced by upstream swagger: **No**.
   - Regression test pins this: **Yes** — `TestSystemDayOfWeek_UnmarshalJSON_IntForm`,
     `_StringForm`, `_InvalidString`, `_OutOfRangeInt`, `TestWeeklyModel_UnmarshalJSON_DayNameStrings`,
     and `_DayIndexInts` in `model_system_day_of_week_test.go` cover both forms and error
     cases.
   - Action: **preserve**.

### `v24/api/keyfactor/v1/model_templates_template_retrieval_response.go` + `model_templates_template_update_request.go`

10. **`af6340b`** — adds `CertificateCleanupEnabled NullableBool`, `TimeAfterExpiration
    NullableInt32`, `TimeAfterExpirationUnits
    *CSSCMSDataModelEnumsCertificateCleanupTimeUnits`, `DeleteWithArchivedKey NullableBool`,
    and (retrieval response only) `Manageability NullableInt32` fields, with
    generator-style getters/setters/Has*/`ToMap` entries. `Manageability` was present in
    actual JSON responses but missing from the Go struct, so it always deserialized as the
    zero value before this fix.
    - Reproduced by upstream swagger: **No**.
    - Regression test pins this: **No** — no test in this repo references
      `TemplatesTemplateRetrievalResponse`, `TemplatesTemplateUpdateRequest`, or
      `Manageability`.
    - Action: **preserve**.

---

## v25 (partial — not a full audit)

The v25 module (baseline `536f3e2`) independently carries its own OAuth-field-restoration
(`0c1df4d`) and port-443-guard (`b374f3d`) hand-edits, predating this document and this
branch. Those have **not** been audited against this document's conventions and are
deliberately **not** claimed as covered here — see "Scope" above.

The one v25 edit made as part of this change set:

1. **`db2be75`** — *fix(v25): plumb Server.ClientTimeout into rebuilt auth config* — the same
   edit as v24 entry #4 above, applied to `v25/api/keyfactor/v1/client.go` and
   `v25/api/keyfactor/v2/client.go`. Required bumping v25's `keyfactor-auth-client-go`
   dependency from `v1.3.0` to `v1.6.0-rc.2` (the first version with `Server.ClientTimeout`);
   verified no API compatibility breaks (`go mod tidy && go build ./...` clean).
   - Reproduced by upstream swagger: **No**.
   - Regression test pins this: **Yes** — `TestBuildHttpClientV2_ClientTimeoutPropagation`,
     ported verbatim from v24, in both v25 v1 and v2 `client_test.go`.
   - Action: **preserve**.

2. Also ported to v25 as part of this change set: `TestPrepareRequest_Port443Guard` (see v24
   entry #3), protecting v25's pre-existing `b374f3d` port-443 guard, which had the same
   "unprotected hand-edit" gap as v24's copy.

**Not fixed in this change set:** the repo-root `api/keyfactor/` module and the `v2/`
(`github.com/Keyfactor/keyfactor-go-client-sdk/v2`) module contain the identical
`buildHttpClientV2` `ClientTimeout`-drop bug. A dependency bump to `keyfactor-auth-client-go
v1.6.0-rc.2` was verified to build cleanly in both (`go mod tidy && go build ./...`,
`GOWORK=off`), so the bug is technically fixable the same way. It was deliberately left
unfixed here because: (a) neither module has a single `*_test.go` file today — fixing "with
equivalent tests" per the standard set by v24/v25 would mean authoring first-ever test
scaffolding for these modules, a materially larger and separate undertaking than the
one-line v24/v25 fix; (b) both appear to be legacy/superseded major-API-version modules (the
README states v25 is "the latest available SDK"; `CHANGELOG.md` has no recent root/v2-module
entries); (c) the known primary consumer (terraform-provider-keyfactor) pins only the v24
module. If an active consumer of the root or `v2` module surfaces, file a follow-up issue
rather than assuming this gap is safe to ignore indefinitely.
