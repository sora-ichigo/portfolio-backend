.PHONY: sam-build deploy migrate-create

S3_BUCKET=igsr5-portfolio-api-lambda-code
OUTPUT_FILE=terraform/template.yaml

# sam
sam-build:
	sam build
package: sam-build
	sam package --s3-bucket $(S3_BUCKET) --output-template-file $(OUTPUT_FILE)
	cd terraform && terraform apply -auto-approve

# go
migrate-create:
	./bin/migrate create -ext sql -dir ./migrations "$(T)"
gen:
	go generate ./...
