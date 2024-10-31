## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |
| <a name="requirement_github"></a> [github](#requirement\_github) | >=6.2 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_github"></a> [github](#provider\_github) | 6.3.1 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_keyfactor_github_test_environment_12_3_0_kc"></a> [keyfactor\_github\_test\_environment\_12\_3\_0\_kc](#module\_keyfactor\_github\_test\_environment\_12\_3\_0\_kc) | git::ssh://git@github.com/Keyfactor/terraform-module-keyfactor-github-test-environment-ad.git | main |
| <a name="module_keyfactor_github_test_environment_ad_10_5_0"></a> [keyfactor\_github\_test\_environment\_ad\_10\_5\_0](#module\_keyfactor\_github\_test\_environment\_ad\_10\_5\_0) | git::ssh://git@github.com/Keyfactor/terraform-module-keyfactor-github-test-environment-ad.git | main |

## Resources

| Name | Type |
|------|------|
| [github_repository.repo](https://registry.terraform.io/providers/integrations/github/latest/docs/data-sources/repository) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_keyfactor_auth_token_url_12_3_0_KC"></a> [keyfactor\_auth\_token\_url\_12\_3\_0\_KC](#input\_keyfactor\_auth\_token\_url\_12\_3\_0\_KC) | The hostname of the KeyCloak instance to authenticate to for a Keyfactor Command access token | `string` | `"https://int-oidc-lab.eastus2.cloudapp.azure.com:8444/realms/Keyfactor/protocol/openid-connect/token"` | no |
| <a name="input_keyfactor_client_id_12_3_0"></a> [keyfactor\_client\_id\_12\_3\_0](#input\_keyfactor\_client\_id\_12\_3\_0) | The client ID to authenticate with the Keyfactor instance using Keycloak client credentials | `string` | n/a | yes |
| <a name="input_keyfactor_client_secret_12_3_0"></a> [keyfactor\_client\_secret\_12\_3\_0](#input\_keyfactor\_client\_secret\_12\_3\_0) | The client secret to authenticate with the Keyfactor instance using Keycloak client credentials | `string` | n/a | yes |
| <a name="input_keyfactor_hostname_10_5_0"></a> [keyfactor\_hostname\_10\_5\_0](#input\_keyfactor\_hostname\_10\_5\_0) | The hostname of the Keyfactor instance | `string` | `"integrations1050-lab.kfdelivery.com"` | no |
| <a name="input_keyfactor_hostname_12_3_0_KC"></a> [keyfactor\_hostname\_12\_3\_0\_KC](#input\_keyfactor\_hostname\_12\_3\_0\_KC) | The hostname of the Keyfactor instance | `string` | `"int-oidc-lab.eastus2.cloudapp.azure.com"` | no |
| <a name="input_keyfactor_password_10_5_0"></a> [keyfactor\_password\_10\_5\_0](#input\_keyfactor\_password\_10\_5\_0) | The password to authenticate with the Keyfactor instance | `string` | n/a | yes |
| <a name="input_keyfactor_username_10_5_0"></a> [keyfactor\_username\_10\_5\_0](#input\_keyfactor\_username\_10\_5\_0) | The username to authenticate with the Keyfactor instance | `string` | n/a | yes |

## Outputs

No outputs.
