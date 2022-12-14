
.PHONY: go-tidy
go-tidy: ## tidy all go modules
	$(call print-target)
	find . -mindepth 2 -name 'go.mod' -execdir go mod tidy -v \;

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: tests
tests: ## run tests of all submodules
	$(call print-target)
	go test -race -timeout 30s -v -coverprofile=coverage.txt $(shell dirname `find . -mindepth 2 -name 'go.mod'`)

.PHONE: workspace
workspace: ## init go workspace 
	go work init $(shell dirname -mindepth `find -mindepth 2 . -name 'go.mod'`)

define print-target
	@printf "Executing target: \033[36m$@\033[0m\n"
endef