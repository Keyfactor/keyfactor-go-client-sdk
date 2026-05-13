# Swagger Gaps — Notes for Engineering

This document captures discrepancies found between the `swagger/Keyfactor-Command-v25-v{1,2}.swagger.json` files and the official Keyfactor Command v25.5.1 API documentation at <https://software.keyfactor.com/Core-OnPrem/v25.5.1/Content/WebAPI/KeyfactorAPI/>.

Most findings are **already resolved** in the swagger files vs. their prior single-file ancestor (`Keyfactor-Command-v10.swagger.yaml`). One actual gap was found and patched in place. Two cross-source disagreements were noticed and are flagged below for next-release consideration.

## Patches applied (2)

### 1. Add `schema: {type: string}` to 475 malformed header parameters

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

### 2. `CSS.CMS.Core.Enums.EnrollmentType` enum: value `4` → `3`

- **File**: `swagger/Keyfactor-Command-v25-v1.swagger.json`
- **Schema**: `components.schemas["CSS.CMS.Core.Enums.EnrollmentType"]`
- **Before**: `enum: [0, 1, 2, 4]`
- **After**: `enum: [0, 1, 2, 3]`

**Rationale**: The Keyfactor docs (`CertificateAuthorityPOST.htm`, "AllowedEnrollmentTypes Values" table) define this enum as a bitmask:

| Value | Description |
|---|---|
| 0 | No enrollment enabled |
| 1 | PFX Enrollment |
| 2 | CSR Enrollment |
| 3 | PFX and CSR Enrollment (= 1 \| 2) |

The swagger's `4` is incorrect — it has no documented meaning. The maximum legal value given only PFX (1) and CSR (2) bits is `3`.

Downstream effect: regenerated SDK code will declare `CSSCMSCOREENUMSENROLLMENTTYPE__3` instead of `CSSCMSCOREENUMSENROLLMENTTYPE__4`. Any consumer code still referencing the (incorrect) `4` constant will need to update — but no real Keyfactor Command instance returns `4` for this field, so behavior in practice should be unchanged.

**Recommended action for engineering**: confirm the swagger source-of-truth has `[0, 1, 2, 3]` going forward and that no internal tooling depends on `4`.

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
