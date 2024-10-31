module "keyfactor_github_test_environment_ad_10_5_0" {
  source = "git::ssh://git@github.com/Keyfactor/terraform-module-keyfactor-github-test-environment-ad.git?ref=main"

  gh_environment_name = "KFC_10_5_0"
  gh_repo_name        = data.github_repository.repo.name
  keyfactor_hostname  = var.keyfactor_hostname_10_5_0
  keyfactor_username  = var.keyfactor_username_10_5_0
  keyfactor_password  = var.keyfactor_password_10_5_0
  keyfactor_config_file = base64encode(file("${path.module}/command_config.json"))
}

# module "keyfactor_github_test_environment_11_5_0_kc" {
#   source = "git::ssh://git@github.com/Keyfactor/terraform-module-keyfactor-github-test-environment-kc.git?ref=main"
#
#   gh_environment_name       = "KFC_11_5_0_KC"
#   gh_repo_name              = data.github_repository.repo.name
#   keyfactor_hostname        = var.keyfactor_hostname_11_5_0_KC
#   keyfactor_client_id       = var.keyfactor_client_id_11_5_0
#   keyfactor_client_secret   = var.keyfactor_client_secret_11_5_0
#   keyfactor_auth_hostname   = var.keyfactor_auth_hostname_11_5_0_KC
#   keyfactor_tls_skip_verify = true
# }

module "keyfactor_github_test_environment_12_3_0_kc" {
  source = "git::ssh://git@github.com/Keyfactor/terraform-module-keyfactor-github-test-environment-ad.git?ref=main"

  gh_environment_name       = "KFC_12_3_0_KC"
  gh_repo_name              = data.github_repository.repo.name
  keyfactor_hostname        = var.keyfactor_hostname_12_3_0_KC
  keyfactor_auth_token_url  = var.keyfactor_auth_token_url_12_3_0_KC
  keyfactor_client_id       = var.keyfactor_client_id_12_3_0
  keyfactor_client_secret   = var.keyfactor_client_secret_12_3_0
  keyfactor_tls_skip_verify = true
  keyfactor_config_file = base64encode(file("${path.module}/command_config.json"))
}
