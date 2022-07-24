BIN_DIR := ${PWD}/bin
export PATH := ${BIN_DIR}:${PATH}

# sam
S3_BUCKET=igsr5-portfolio-api-lambda-code
OUTPUT_FILE=terraform/template.yaml

.PHONY: sam-build
sam-build:
	sam build
.PHONY: sam-local
sam-local:
	sam local start-api --docker-network portfolio-backend_network --env-vars env-local.json
.PHONY: deploy
package: sam-build
	sam package --s3-bucket $(S3_BUCKET) --output-template-file $(OUTPUT_FILE)
	cd terraform && terraform apply -auto-approve

# migrate
DB_DRIVER=mysql://
DSN=root:root@tcp(localhost:3306)/portfolio?parseTime=true
STEP=1

.PHONY: migrate-create
migrate-create: tools
	./bin/migrate create -ext sql -dir ./migrations "$(T)"
.PHONY: migrate
migrate: tools
	./bin/migrate -path migrations/ -database "$(DB_DRIVER)$(DSN)" up
.PHONY: rollback
rollback: tools
	./bin/migrate -path migrations/ -database "$(DB_DRIVER)$(DSN)" down "$(STEP)"
.PHONY: migrate-force
migrate-force: tools
	./bin/migrate -path migrations/ -database "$(DB_DRIVER)$(DSN)" force "$(VERSION)"

# go
.PHONY: tools
tools:
	go generate ./tools.go
.PHONY: gen
gen:
	go generate ./gen.go
.PHONY: test
test:
	DSN="$(DSN)" go test ./...
