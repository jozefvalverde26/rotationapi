.PHONY: install-tools

install-tools: ## Install external tools
	@bash scripts/install-tools.sh

install-deps: ## Prefetch deps to ensure required versions are downloaded
	@go mod tidy
	@go mod verify