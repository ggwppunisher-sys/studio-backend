.PHONY: generate-oapi run migrations-new migrations-up migrations-down migrations-status

# Common command variables
OAPI_CMD := go tool oapi-codegen
GOOSE_CMD := go tool goose -v -env $(or $(envfile),cluster/example.env)

generate-oapi:
	@echo 'Generating http server...'
	$(OAPI_CMD) -config oapi/oapi.server.cfg.yaml oapi/openapi.yaml
	@echo 'Generating http client...'
	$(OAPI_CMD) -config oapi/oapi.client.cfg.yaml oapi/openapi.yaml

run:
	@go run cmd/main.go --env $(or $(envfile),cluster/example.env)

# Create new db migration file
# Usage: make migrations-new filename=create_users_table envfile=example.env
migrations-new:
	$(GOOSE_CMD) create $(filename) sql

# Apply all pending migrations
# Usage: make migrations-up envfile=example.env
migrations-up:
	@echo 'Applying migrations...'
	$(GOOSE_CMD) up -v

# Down one last migration
# Usage: make migrations-down envfile=example.env
migrations-down:
	$(GOOSE_CMD) down -v

# Check migrations status
# Usage: make migrations-status envfile=example.env
migrations-status:
	$(GOOSE_CMD) status
