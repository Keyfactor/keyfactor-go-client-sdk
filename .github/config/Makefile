.DEFAULT_GOAL := help

##@ Utility
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

deps: ## Install deps for macos
	@brew install pre-commit tflint terraform terraform-docs

docs: ## Run terraform-docs to update module docs.
	@terraform-docs markdown . > MODULE.MD
	@terraform-docs markdown table --output-file README.md --output-mode inject .

lint: ## Run tflint
	@tflint

validate: ## Run terraform validate
	@terraform init --upgrade
	@terraform validate

precommit/add: ## Install pre-commit hook
	@pre-commit install

precommit/remove: ## Uninstall pre-commit hook
	@pre-commit uninstall

