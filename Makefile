
# sam
S3_BUCKET=igsr5-portfolio-api-lambda-code
OUTPUT_FILE=terraform/template.yaml

.PHONY: sam-build
sam-build:
	sam build
.PHONY: deploy
package: sam-build
	sam package --s3-bucket $(S3_BUCKET) --output-template-file $(OUTPUT_FILE)
	cd terraform && terraform apply -auto-approve

# migrate
DSN=mysql://root:root@tcp(db:3306)/portfolio
STEP=1

.PHONY: migrate-create
migrate-create: gen
	./bin/migrate create -ext sql -dir ./migrations "$(T)"
.PHONY: migrate
migrate: gen
	./bin/migrate -path migrations/ -database "$(DSN)" up
.PHONY: rollback
rollback: gen
	./bin/migrate -path migrations/ -database "$(DSN)" down "$(STEP)"
.PHONY: migrate-force
migrate-force: gen
	./bin/migrate -path migrations/ -database "$(DSN)" force "$(VERSION)"

# go
.PHONY: gen
gen:
	go generate ./...
