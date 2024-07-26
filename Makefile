.PHONY: generate
generate:
	docker compose up
	$(format)

define format
	@go fmt ./... 
	@go run golang.org/x/tools/cmd/goimports -w ./ 
	@go run mvdan.cc/gofumpt -l -w .
	@go mod tidy
endef