terraform {
  required_version = ">= 1.0"
  required_providers {
    github = {
      source = "integrations/github"
      version = ">=6.2"
    }
  }
  backend "azurerm" {
    resource_group_name  = "integrations-infra"
    storage_account_name = "integrationstfstate"
    container_name       = "tfstate"
    key                  = "github/repos/keyfactor-auth-client-go.tfstate"
  }
}

provider "github" {
  # Configuration options
  owner = "Keyfactor"
}