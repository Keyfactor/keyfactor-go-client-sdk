# Swagger Gaps — Notes for Engineering

This document captures discrepancies found between the `swagger/Keyfactor-Command-v25-v{1,2}.swagger.json` files and the official Keyfactor Command v25.5.1 API documentation at <https://software.keyfactor.com/Core-OnPrem/v25.5.1/Content/WebAPI/KeyfactorAPI/>.

Most findings are **already resolved** in the swagger files vs. their prior single-file ancestor (`Keyfactor-Command-v10.swagger.yaml`). One actual gap was found and patched in place. Two cross-source disagreements were noticed and are flagged below for next-release consideration.

## Patches applied (3)

### 1. Strip `Keyfactor.Web.KeyfactorApi.Models.` namespace prefix from schema names

This patch is **applied at regen-time only** (via `scripts/strip-schema-namespace.py`) — the source swagger files in `swagger/` are not modified.

**Issue**: every schema in the upstream v25 swagger now carries a fully-qualified namespace prefix `Keyfactor.Web.KeyfactorApi.Models.` — for example:

- `Keyfactor.Web.KeyfactorApi.Models.Templates.TemplateRetrievalResponse`
- `Keyfactor.Web.KeyfactorApi.Models.CertificateAuthorities.CertificateAuthorityRequest`
- `Keyfactor.Web.KeyfactorApi.Models.PAMProviderResponse`

The swagger that produced the originally-released v25 SDK (`536f3e2`) used trailing names only:

- `Templates.TemplateRetrievalResponse`
- `CertificateAuthorities.CertificateAuthorityRequest`
- `PAMProviderResponse`

Since `openapi-generator` translates schema names directly into Go file and type names, every model file and every type identifier in the regenerated v25 code would change. This breaks:
- `v25/client.go` (hand-written wrapper that references hundreds of model types)
- Any downstream consumer (`kfutil`, `terraform-provider-keyfactor`) that imports v25 types
- Every published SDK consumer's compilation

Counts on the v25 swagger:
- v1: 345 schemas + 1491 `$ref` pointers carry the prefix
- v2: 39 schemas + 122 `$ref` pointers

**Workaround in this repo**: `scripts/strip-schema-namespace.py` rewrites every schema name and `$ref` pointer containing the prefix to drop it, before the generator runs. Output naming then matches the historical v25 SDK convention. Source swagger files in `swagger/` are untouched so version-controlled history stays stable.

**Recommended action for engineering**: decide canonically which convention v25+ should use, then make the source swagger generation emit consistently. Options:

1. **Restore the short names** (matches existing v25 SDK and downstream consumers) — engineering can patch the upstream swagger generation to drop the `Keyfactor.Web.KeyfactorApi.Models.` namespace from schema names. This keeps the SDK API stable for consumers.
2. **Keep the long names as canonical** (matches the new swagger output as-is) — engineering documents this as a breaking SDK change; we remove the preprocessor and propagate the rename through `v25/client.go` and downstream consumers.

Other non-`Keyfactor.Web.KeyfactorApi.Models.` Keyfactor namespaces (e.g., `Keyfactor.Common.Scheduling.*`, `Keyfactor.Platform.Extensions.Enums.*`) are already short and need no change.

### 2. Add `schema: {type: string}` to 475 malformed header parameters

- **Files**:
  - `swagger/Keyfactor-Command-v25-v1.swagger.json` — 451 patches (447 × `x-keyfactor-api-version`, 4 × `x-certificateformat`)
  - `swagger/Keyfactor-Command-v25-v2.swagger.json` — 24 patches (23 × `x-keyfactor-api-version`, 1 × `x-certificateformat`)

**Issue**: every `x-keyfactor-api-version` and `x-certificateformat` header parameter in both swagger files is declared as `{name, in, description, [required], [example]}` with **no `schema` field and no `content` field**. OpenAPI 3.x requires non-body parameters to have one or the other. Comparison with a sibling well-formed parameter from the same swagger:

```jsonc
// CORRECTLY formed (sibling)
{
  "name": "x-keyfactor-requested-with",
  "in": "header",
  "description": "Type of the request [XMLHttpRequest, APIClient]",
  "required": true,
  "schema": { "type": "string" },          // <-- present
  "example": "APIClient"
}

// MALFORMED (in upstream swagger)
{
  "name": "x-keyfactor-api-version",
  "in": "header",
  "description": "Desired version of the api, if not provided defaults to v1",
                                            // <-- schema / content missing
  "example": "1.0"
}
```

**Effect**: openapi-generator-cli 6.3.0 fails with `SpecValidationException` listing 451 errors of the form `parameters.[x-keyfactor-api-version].content is missing` — and refuses to generate code. (The validator misreports it as "content missing"; both `schema` and `content` would satisfy the requirement.)

**Patch**: add `"schema": { "type": "string" }` to every malformed header parameter, inserted before `"example"` to match the ordering of well-formed siblings. Diff is mechanical: ~13KB additions of identical `schema` blocks, plus minor empty-array whitespace normalization (`"Bearer": [ ]` → `"Bearer": []`) introduced by the JSON re-serialization.

**Recommended action for engineering**: fix the upstream swagger generation so these two header parameters emit a `schema` field at the source. The fix is mechanical and presumably driven by the OpenAPI/JSON schema spec annotation system.

### 3. `CSS.CMS.Core.Enums.EnrollmentType` enum: `[0, 1, 2, 4]` → `[0, 1, 2, 3, 4, 5, 6, 7]`

- **File**: `swagger/Keyfactor-Command-v25-v1.swagger.json`
- **Schema**: `components.schemas["CSS.CMS.Core.Enums.EnrollmentType"]`
- **Before**: `enum: [0, 1, 2, 4]`
- **After**: `enum: [0, 1, 2, 3, 4, 5, 6, 7]`

**Rationale**: Originally this was patched to `[0, 1, 2, 3]` based on the docs (`CertificateAuthorityPOST.htm`, "AllowedEnrollmentTypes Values" table showing 0=None, 1=PFX, 2=CSR, 3=PFX+CSR). However, the canonical generation script at `Keyfactor/API-definitions` (`go/command/openapi-generate.sh`) applies the broader patch `[0, 1, 2, 3, 4, 5, 6, 7]` with this comment:

> `# Hotfix -- update the enums array of enrollment types. Zendesk tix 139784`

Engineering has Zendesk ticket evidence that the API returns enrollment-type values beyond `3` in some contexts. The full 0–7 range covers these observed values. The public docs are incomplete — they cover the basic PFX/CSR bitmask but omit the additional values.

**Recommended action for engineering**:
1. Document what enrollment-type values `4, 5, 6, 7` actually mean (the public CertificateAuthorityPOST.htm only documents 0–3).
2. Decide whether the canonical Command swagger should emit the full `[0..7]` range or whether some values are deprecated/internal.
3. Cross-reference with Zendesk 139784 for the original incident report.

---

## Hand-edits we found but DID NOT patch

These were enum values added by hand to the generated Go code (`v25/api/keyfactor/v1/model_css_cms_core_enums_enrollment_type.go`) in commit `2d5c09e`:

- `CSSCMSCOREENUMSENROLLMENTTYPE__5`
- `CSSCMSCOREENUMSENROLLMENTTYPE__6`
- `CSSCMSCOREENUMSENROLLMENTTYPE__7`

**No documentation confirms these values exist.** The documented enum maxes at `3` (PFX | CSR). Values `5, 6, 7` would require additional bits beyond PFX and CSR — but the docs and the swagger both reflect a 2-bit space.

**Recommended action for engineering**: review the source of these hand-edited values. If Command actually returns `5, 6, 7` in any context, the documentation and swagger are both incomplete. If not, the hand-edit should be removed and the regression intent re-evaluated (`2d5c09e: fix(enrollment type): Fix the allowed enum values for enrollment patterns`).

---

## Cross-source disagreements (worth resolving in next-release docs)

### Spelling: `DeleteWithArchivedKey` vs `DeleteWithArchiveKey`

- **Swagger** (`Keyfactor-Command-v25-v1.swagger.json`): `DeleteWithArchivedKey` (with 'd' — past participle). Appears in `CertificateAuthorityRequest`, `CertificateAuthorityResponse`, `TemplateRetrievalResponse`, `TemplateUpdateRequest`.
- **Docs site** (`CertificateAuthorityPOST.htm`, `TemplatesGetID.htm`, `TemplatesPUT.htm`): `DeleteWithArchiveKey` (no 'd'). Documented verbatim in all three pages.

These are inconsistent. The swagger appears to be canonical (matches actual Command API field names used by consumers), so the docs site likely has a typo. **Recommended action**: update the docs site to use `DeleteWithArchivedKey`.

### `Manageability` field on templates

- **Swagger**: `TemplateRetrievalResponse.Manageability` is present, typed as `Keyfactor.Platform.Extensions.Enums.TemplateDetailsManageability` (enum: `[0, 1, 2]`).
- **Docs site** (`TemplatesGet.htm`, `TemplatesGetID.htm`, `TemplatesPUT.htm`): field is **not mentioned** in any documented template endpoint.

The swagger seems to be ahead of the docs here. Either the field is undocumented but real (Command returns it but customers wouldn't know to expect it), or the docs are missing it. **Recommended action**: add `Manageability` to the documented Templates GET-by-id response with the three enum values and their meanings.

### CertificateStoreType `Property.Type` and `EntryParameter.Type` enums declared as bare ints

- **Swagger** (`Keyfactor-Command-v25-v1.swagger.json`):
  - `CSS.CMS.Core.Enums.CertificateStoreTypePropertyType` — `{"type":"integer","format":"int32","enum":[0,1,2,3]}`
  - `CSS.CMS.Core.Enums.CertStoreEntryParameterType` — `{"type":"integer","format":"int32","enum":[0,1,2,3]}`
  - No `x-enum-varnames` / `x-enum-descriptions` annotations. The generated v25 Go SDK therefore exposes `*int32` typed enums with constants named only `__0`, `__1`, `__2`, `__3` and a `Parse(string)` method whose `stringsToEnum` map is empty (so unmarshaling always errors).
- **Actual wire format used by Keyfactor Command**: STRING NAMES, not integers. The deprecated `keyfactor-go-client/v3` package declares both fields as plain `string` (`v3/api/store_type_models.go`: `Type string \`json:"Type"\``), and `kfutil`'s `cmd/store_types.json` ships 70 store-type templates with values like `"String"`, `"Bool"`, `"Secret"`, `"MultipleChoice"`. These templates POST cleanly to the Keyfactor server today.

The swagger and SDK are wrong for both directions:
- **POST/PUT requests**: server accepts string names; the typed SDK body forces ints and would fail to serialize any existing user manifest.
- **GET responses**: empirically the server returns string names on the wire (otherwise the v3 client would have broken long ago); the SDK's `UnmarshalJSON` for these enums rejects strings with `"%+v is not a valid …"`.

**Recommended actions** (any one of these unblocks downstream consumers):
1. Re-declare the schemas as `{"type":"string","enum":["String","MultipleChoice","Bool","Secret"]}` for the property type, and the equivalent set for the entry-parameter type. This matches the documented and observed wire shape and is what the Command server already accepts.
2. If the integer form is genuinely the canonical wire shape (and the server only happens to accept strings as a legacy shim), add `x-enum-varnames` to lock in the int↔name mapping authoritatively so consumers don't have to guess. Then fix the SDK generator's broken `Parse()` method so it actually populates `stringsToEnum`.
3. Either way, document the four enum values explicitly on the docs site (`AddCertificateStoreType.htm` / store-type POST/PUT pages) — they are not currently named anywhere consumer-facing.

**Workaround in kfutil while this is unresolved**: kfutil's storeTypes migration uses the SDK for transport (auth + URL config + headers) but defines a local request/response struct that mirrors v3's `string`-typed `Type` field, preserving wire compatibility with all existing user manifests. This is documented in `kfutil/cmd/storeTypes.go` so it can be reverted once the SDK exposes a working enum.

---

## What was NOT a gap (sanity checks)

These were on the original Phase 0 gap list but turned out to already be in the new swagger — no patch needed:

| Field | Schema in swagger | Confirmed in docs |
|---|---|---|
| `UseForEnrollment` on CA | `…CertificateAuthorityRequest.UseForEnrollment` | ✅ `CertificateAuthorityPOST.htm` |
| `CertificateCleanupEnabled` on CA | `…CertificateAuthorityRequest.CertificateCleanupEnabled` | ✅ same page |
| `TimeAfterExpiration` on CA | same | ✅ |
| `TimeAfterExpirationUnits` on CA | typed `CertificateCleanupTimeUnits` enum (vs docs' string `days/weeks/months`) | ✅ — swagger more precise than docs |
| `DeleteWithArchivedKey` on CA | same | ✅ (modulo spelling discrepancy above) |
| `CertificateCleanupEnabled` on templates | `TemplateRetrievalResponse` + `TemplateUpdateRequest` | ✅ `TemplatesGetID.htm`, `TemplatesPUT.htm` |
| `TimeAfterExpiration` on templates | same | ✅ |
| `TimeAfterExpirationUnits` on templates | same | ✅ |
| `DeleteWithArchivedKey` on templates | same | ✅ (modulo spelling) |
| `KeyRetentionPolicy` enum values `[0, 1, 2, 3]` | already correct | ✅ |
| `CertificateCleanupTimeUnits` enum `[0, 1, 2]` | typed enum, values map to `days/weeks/months` | ✅ |

All of these existed in the **new** swagger but **not** in the prior `Keyfactor-Command-v10.swagger.yaml`. Refreshing the swagger has therefore obsoleted most of the post-generation hand-edits in `v25/api/keyfactor/v1/`. The only remaining hand-edits that still need preservation through regen are **pure code-level** fixes in `client.go` (port-443 handling, OAuth access-token flow, restored OAuth config fields) and the regression tests added in commit `229db7d`.

---

## Source documentation pages consulted

- `https://software.keyfactor.com/Core-OnPrem/v25.5.1/Content/WebAPI/KeyfactorAPI/CertificateAuthorityPOST.htm`
- `https://software.keyfactor.com/Core-OnPrem/v25.5.1/Content/WebAPI/KeyfactorAPI/TemplatesGet.htm`
- `https://software.keyfactor.com/Core-OnPrem/v25.5.1/Content/WebAPI/KeyfactorAPI/TemplatesGetID.htm`
- `https://software.keyfactor.com/Core-OnPrem/v25.5.1/Content/WebAPI/KeyfactorAPI/TemplatesPUT.htm`
- `https://software.keyfactor.com/Core-OnPrem/v25.5.1/Content/WebAPI/KeyfactorAPI/Enrollment-Patterns.htm` (index only — endpoint pages not yet reviewed)
- `https://software.keyfactor.com/Core-OnPrem/v25.5.1/Content/WebAPI/KeyfactorAPI/Enrollment.htm` (index only)
- `https://software.keyfactor.com/Core-OnPrem/v25.5.1/Content/WebAPI/KeyfactorAPI/CertificateAuthority.htm` (index only)
