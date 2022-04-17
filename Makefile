.PHONY: build deploy

S3_BUCKET=s3://igsr5-portfolio-api-lambda

build:
	sam build
deploy: build
	# sam package --s3-bucket $(S3_BUCKET) --output-template-file ./terraform/cloudformation
	cd terraform && terraform apply
