.PHONY: build deploy

S3_BUCKET=igsr5-portfolio-api-lambda-code
OUTPUT_FILE=terraform/template.yaml

build:
	sam build
package: build
	sam package --s3-bucket $(S3_BUCKET) --output-template-file $(OUTPUT_FILE)
	cd terraform && terraform apply -auto-approve
