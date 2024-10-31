# GitHub Test Environment Setup

This code sets up GitHub environments for testing against Keyfactor Command instances that are configured to use
Active Directory or Keycloak for authentication.

## Requirements

1. Terraform >= 1.0
2. GitHub Provider >= 6.2
3. Keyfactor Command instance(s) configured to use Active Directory or Keycloak for authentication
4. AD or Keycloak credentials for authenticating to the Keyfactor Command instance(s)
5. A GitHub token with access and permissions to the repository where the environments will be created

## Adding a new environment

Modify the `environments.tf` file to include the new environment module. The module should be named appropriately.
Example:

### Active Directory Environment

```hcl
module "keyfactor_github_test_environment_ad_10_5_0" {
  source = "git::ssh://git@github.com/Keyfactor/terraform-module-keyfactor-github-test-environment-ad.git?ref=main"

  gh_environment_name = "KFC_10_5_0" # Keyfactor Command 10.5.0 environment using Active Directory(/Basic Auth)
  gh_repo_name       = data.github_repository.repo.name
  keyfactor_hostname = var.keyfactor_hostname_10_5_0
  keyfactor_username = var.keyfactor_username_10_5_0
  keyfactor_password = var.keyfactor_password_10_5_0
}
```

### oAuth Client Environment

```hcl
module "keyfactor_github_test_environment_12_3_0_kc" {
  source = "git::ssh://git@github.com/Keyfactor/terraform-module-keyfactor-github-test-environment-kc.git?ref=main"

  gh_environment_name = "KFC_12_3_0_KC" # Keyfactor Command 12.3.0 environment using Keycloak
  gh_repo_name              = data.github_repository.repo.name
  keyfactor_hostname        = var.keyfactor_hostname_12_3_0_KC
  keyfactor_auth_token_url  = var.keyfactor_auth_token_url_12_3_0_KC
  keyfactor_client_id       = var.keyfactor_client_id_12_3_0
  keyfactor_client_secret   = var.keyfactor_client_secret_12_3_0
  keyfactor_tls_skip_verify = true
}
```

<!-- BEGIN_TF_DOCS -->

## Requirements

| Name                                                                      | Version |
|---------------------------------------------------------------------------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0  |
| <a name="requirement_github"></a> [github](#requirement\_github)          | >=6.2   |

## Providers

| Name                                                       | Version |
|------------------------------------------------------------|---------|
| <a name="provider_github"></a> [github](#provider\_github) | 6.3.1   |

## Modules

| Name                                                                                                                                                                                | Source                                                                                        | Version |
|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|---------|
| <a name="module_keyfactor_github_test_environment_12_3_0_kc"></a> [keyfactor\_github\_test\_environment\_12\_3\_0\_kc](#module\_keyfactor\_github\_test\_environment\_12\_3\_0\_kc) | git::ssh://git@github.com/Keyfactor/terraform-module-keyfactor-github-test-environment-ad.git | main    |
| <a name="module_keyfactor_github_test_environment_ad_10_5_0"></a> [keyfactor\_github\_test\_environment\_ad\_10\_5\_0](#module\_keyfactor\_github\_test\_environment\_ad\_10\_5\_0) | git::ssh://git@github.com/Keyfactor/terraform-module-keyfactor-github-test-environment-ad.git | main    |

## Resources

| Name                                                                                                                      | Type        |
|---------------------------------------------------------------------------------------------------------------------------|-------------|
| [github_repository.repo](https://registry.terraform.io/providers/integrations/github/latest/docs/data-sources/repository) | data source |

## Inputs

| Name                                                                                                                                                   | Description                                                                                     | Type     | Default                                                                                                 | Required |
|--------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|----------|---------------------------------------------------------------------------------------------------------|:--------:|
| <a name="input_keyfactor_auth_token_url_12_3_0_KC"></a> [keyfactor\_auth\_token\_url\_12\_3\_0\_KC](#input\_keyfactor\_auth\_token\_url\_12\_3\_0\_KC) | The hostname of the KeyCloak instance to authenticate to for a Keyfactor Command access token   | `string` | `"https://int-oidc-lab.eastus2.cloudapp.azure.com:8444/realms/Keyfactor/protocol/openid-connect/token"` |    no    |
| <a name="input_keyfactor_client_id_12_3_0"></a> [keyfactor\_client\_id\_12\_3\_0](#input\_keyfactor\_client\_id\_12\_3\_0)                             | The client ID to authenticate with the Keyfactor instance using Keycloak client credentials     | `string` | n/a                                                                                                     |   yes    |
| <a name="input_keyfactor_client_secret_12_3_0"></a> [keyfactor\_client\_secret\_12\_3\_0](#input\_keyfactor\_client\_secret\_12\_3\_0)                 | The client secret to authenticate with the Keyfactor instance using Keycloak client credentials | `string` | n/a                                                                                                     |   yes    |
| <a name="input_keyfactor_hostname_10_5_0"></a> [keyfactor\_hostname\_10\_5\_0](#input\_keyfactor\_hostname\_10\_5\_0)                                  | The hostname of the Keyfactor instance                                                          | `string` | `"integrations1050-lab.kfdelivery.com"`                                                                 |    no    |
| <a name="input_keyfactor_hostname_12_3_0_KC"></a> [keyfactor\_hostname\_12\_3\_0\_KC](#input\_keyfactor\_hostname\_12\_3\_0\_KC)                       | The hostname of the Keyfactor instance                                                          | `string` | `"int-oidc-lab.eastus2.cloudapp.azure.com"`                                                             |    no    |
| <a name="input_keyfactor_password_10_5_0"></a> [keyfactor\_password\_10\_5\_0](#input\_keyfactor\_password\_10\_5\_0)                                  | The password to authenticate with the Keyfactor instance                                        | `string` | n/a                                                                                                     |   yes    |
| <a name="input_keyfactor_username_10_5_0"></a> [keyfactor\_username\_10\_5\_0](#input\_keyfactor\_username\_10\_5\_0)                                  | The username to authenticate with the Keyfactor instance                                        | `string` | n/a                                                                                                     |   yes    |

## Outputs

No outputs.
<!-- END_TF_DOCS -->