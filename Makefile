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
	sam local start-api --docker-network portfolio-server-api_portfolio
.PHONY: deploy
package: sam-build
	sam package --s3-bucket $(S3_BUCKET) --output-template-file $(OUTPUT_FILE)
	cd terraform && terraform apply -auto-approve

# migrate
DSN=mysql://root:root@tcp(localhost:3306)/portfolio
STEP=1

.PHONY: migrate-create
migrate-create: tools
	./bin/migrate create -ext sql -dir ./migrations "$(T)"
.PHONY: migrate
migrate: tools
	./bin/migrate -path migrations/ -database "$(DSN)" up
.PHONY: rollback
rollback: tools
	./bin/migrate -path migrations/ -database "$(DSN)" down "$(STEP)"
.PHONY: migrate-force
migrate-force: tools
	./bin/migrate -path migrations/ -database "$(DSN)" force "$(VERSION)"

# go
.PHONY: tools
tools:
	go generate ./tools.go
.PHONY: gen
gen:
	go generate ./gen.go
.PHONY: test
test:
	go test -v ./...
