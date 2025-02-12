# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
BINFILE=tarela

help: ## Display this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the code
	mkdir -p ./dist
	go build -o dist/${BINFILE}

test: ## Run all tests
	go test -v ./...

clean: ## Remove the dist directory
	rm -r ./dist