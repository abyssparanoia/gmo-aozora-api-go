.PHONY: generate
generate:
	docker compose up
	$(format)

.PHONY: format
format:
	$(format)

define format
	@go fmt ./... 
	@go run golang.org/x/tools/cmd/goimports -w ./ 
	@go run mvdan.cc/gofumpt -l -w .
	@go run github.com/google/yamlfmt/cmd/yamlfmt .
	@go mod tidy
endef