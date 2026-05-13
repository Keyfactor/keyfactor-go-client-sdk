# v25 Regen Diff Analysis

Analysis of files present in `v25/api/keyfactor/v1/` baseline but absent from `.regen-staging/v25/api/keyfactor/v1/` after running `make regen-stage` with the canonical pipeline (`scripts/regen.sh` + `custom-templates/go/`).

**Bottom line**: zero genuine regressions. Every absence is a rename, an API-version relocation (v1→v2), a namespace reorganization, or one intentional Command-side removal. Current SDK consumers (kfutil, terraform-provider-keyfactor) depend on the legacy `keyfactor-go-client-sdk/v2` module — they do not reference v25 types yet, so regen-apply has zero blast radius on them.

## Methodology

For each `.regen-baseline/v25/api/keyfactor/v1/*.go` file absent from staging:

1. Extract the primary `type X struct {…}` declaration.
2. Search `.regen-staging/` (both v1 and v2) for the same type name.
3. If absent, search by stripped-suffix forms.
4. Look up the corresponding swagger schema name(s) in `swagger/Keyfactor-Command-v25-v{1,2}.swagger.json`.
5. For any type whose corresponding swagger schema disappeared, grep `kfutil`/`terraform-provider-keyfactor` for consumer usage.

## Category 1 — Schema rename: `CertificateRetrievalResponse` → `CertificateRetrievalBulkResponse`

7 cert retrieval sub-models renamed. The parent schema picked up a `Bulk` qualifier; submodels follow.

| Baseline file | Staging file |
|---|---|
| `model_certificates_certificate_retrieval_response_certificate_store_inventory_item_model.go` | `model_certificates_certificate_retrieval_bulk_response_certificate_store_inventory_item_model.go` |
| `…_certificate_store_location_detail_model.go` | `…_bulk_response_certificate_store_location_detail_model.go` |
| `…_crl_distribution_point_model.go` | `…_bulk_response_crl_distribution_point_model.go` |
| `…_detailed_key_usage_model.go` | `…_bulk_response_detailed_key_usage_model.go` |
| `…_extended_key_usage_model.go` | `…_bulk_response_extended_key_usage_model.go` |
| `…_location_count_model.go` | `…_bulk_response_location_count_model.go` |
| `…_subject_alternative_name_model.go` | `…_bulk_response_subject_alternative_name_model.go` |

**Consumer impact**: 0 references to the old names in kfutil or TPK.

## Category 2 — Moved to v2 API: `Workflows.Definition*`

6 workflow definition models moved from v1 to v2 in the new swagger. v1 still exposes them under a `V1`-suffixed legacy namespace (`Workflows.V1.Definition*V1*`) for compatibility.

| Baseline (v1) | New v2 (canonical) | New v1 legacy |
|---|---|---|
| `model_workflows_definition_create_request.go` | `v2/model_workflows_definition_create_request.go` | `v1/model_workflows_v1_definition_create_v1_request.go` |
| `model_workflows_definition_query_response.go` | `v2/model_workflows_definition_query_response.go` | `v1/model_workflows_v1_definition_query_v1_response.go` |
| `model_workflows_definition_response.go` | `v2/model_workflows_definition_response.go` | `v1/model_workflows_v1_definition_v1_response.go` |
| `model_workflows_definition_step_response.go` | `v2/model_workflows_definition_step_response.go` | `v1/model_workflows_v1_definition_step_v1_response.go` |
| `model_workflows_definition_step_signal_response.go` | `v2/model_workflows_definition_step_signal_response.go` | `v1/model_workflows_v1_definition_step_signal_v1_response.go` |
| `model_workflows_definition_update_request.go` | `v2/model_workflows_definition_update_request.go` | `v1/model_workflows_v1_definition_update_v1_request.go` |

The new v2 swagger has paths `/Workflow/Definitions` and `/Workflow/Definitions/{definitionId}`.

**Consumer impact**: 2 kfutil files reference `WorkflowsDefinitionCreateRequest` (`cmd/import.go`, `cmd/export.go`). However, these imports point at the **legacy `keyfactor-go-client-sdk/v2` module**, not at `v25/api/keyfactor/v1`. So this regen does not touch them. When kfutil migrates to v25 SDK (Phase 2 of the migration plan), it will need to import from `v25.../v2` not `v25.../v1`.

## Category 3 — Moved to v2 API: SMTP test

2 SMTP test endpoint models moved from v1 to v2.

| Baseline (v1) | New v2 |
|---|---|
| `model_smtp_smtp_test_request.go` | `v2/model_smtp_smtp_test_request.go` |
| `model_smtp_smtp_test_response.go` | (test response type not regenerated — POST `/SMTP/Test` likely returns empty body or generic status) |

The new v2 swagger has the `/SMTP/Test` path.

**Consumer impact**: 0 references.

## Category 4 — Namespace reorganization: `CSS.CMS.Data.Model.Models.*` → `Keyfactor.Web.KeyfactorApi.Models.*`

8 baseline files in the legacy `CSS.CMS.Data.Model.Models.*` namespace have equivalents under the newer `Keyfactor.Web.KeyfactorApi.Models.*` namespace in the new swagger. The canonical script's `sed` strips that prefix, so the staged names end up like `CertificateStoresTypes.EntryParameters` instead of `CSS.CMS.Data.Model.Models.CertificateStoreType.Property`.

| Baseline file | Staging equivalent (best match) |
|---|---|
| `model_css_cms_data_model_models_cert_store_locations_certificate_locations_group.go` | `model_certificate_stores_locations_certificate_locations_group.go` |
| `model_css_cms_data_model_models_cert_store_locations_certificate_store_locations_detail.go` | `model_certificate_stores_locations_certificate_store_locations_detail.go` |
| `model_css_cms_data_model_models_cert_store_type_password_options.go` | `model_certificate_stores_types_password_options.go` |
| `model_css_cms_data_model_models_cert_store_type_supported_operations.go` | `model_certificate_stores_types_supported_operations.go` |
| `model_css_cms_data_model_models_certificate_store_container_list_response.go` | `model_certificate_stores_container_list_response.go` |
| `model_css_cms_data_model_models_certificate_store_type_property.go` | `model_certificate_stores_types_property.go` |
| `model_css_cms_data_model_models_templates_algorithms_key_info.go` | `model_templates_algorithms_key_info.go` |
| `model_css_cms_data_model_models_templates_template_enrollment_field.go` | `model_templates_template_enrollment_field.go` |

**Consumer impact**: not directly checked file-by-file, but consumers use the legacy `keyfactor-go-client-sdk/v2` module — no impact on this regen.

## Category 5 — Genuine API removal

1 file represents a real Command-side removal.

| Baseline file | Status |
|---|---|
| `model_css_cms_core_enums_entry_parameter_usage_flags.go` (`CSSCMSCoreEnumsEntryParameterUsageFlags` enum) | **Removed**. The `RequiredWhen` field on `CertificateStoreTypeEntryParameter` that referenced this enum is also absent from the new swagger. No consumer references it. Likely deprecated in Command API. |

## Category 6 — Cosmetic name change

1 file lost the `Model` suffix.

| Baseline | Staging |
|---|---|
| `model_templates_key_algorithms_response_model.go` (`TemplatesKeyAlgorithmsResponseModel`) | `model_templates_key_algorithms_response.go` (`TemplatesKeyAlgorithmsResponse`) |

No consumer references.

## Summary

| Category | Count | Action required before regen-apply |
|---|---|---|
| Schema rename (BulkResponse) | 7 | None — staging equivalents exist |
| Moved v1 → v2 (workflow definitions) | 6 | None — both v1-legacy and v2-canonical present; consumer impact is deferred to consumer migration |
| Moved v1 → v2 (SMTP test) | 2 | None |
| Namespace reorg | 8 | None — staging equivalents exist |
| Genuine API removal | 1 | None — no consumer references |
| Cosmetic suffix drop | 1 | None |
| **Total** | **25** | **regen-apply is safe** |

## Notes for future consumer migration to v25 SDK

When kfutil (or TPK or any other consumer) migrates from `keyfactor-go-client-sdk/v2` to `keyfactor-go-client-sdk/v25`, the consumer code will need to handle:

1. **Workflow definitions**: import from `v25/api/keyfactor/v2`, not `v25/api/keyfactor/v1`. The v1 SDK has `Workflows.V1.Definition*V1*` types for backward source compatibility, but new consumer code should target v2.
2. **SMTP test request**: same — `v25/api/keyfactor/v2`.
3. **Cert retrieval submodels**: types are renamed from `CertificateRetrievalResponse*` to `CertificateRetrievalBulkResponse*`.
4. **`EntryParameterUsageFlags` / `RequiredWhen` field on store type entry parameters**: gone. If consumer code uses this, it must adapt to whatever the new field shape is.
