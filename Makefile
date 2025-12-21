.PHONY: generate-oapi run


generate-oapi:
	@echo 'Generating http server...'
	go tool oapi-codegen -config oapi/oapi.server.cfg.yaml oapi/openapi.yaml
	@echo 'Generating http client...'
	go tool oapi-codegen -config oapi/oapi.client.cfg.yaml oapi/openapi.yaml

run:
	@go run cmd/main.go --env example.env
