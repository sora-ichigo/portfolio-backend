resource "aws_s3_bucket" "igsr5-portfolio-api-lambda-code" {
  bucket = "igsr5-portfolio-api-lambda-code"
}

resource "aws_cloudformation_stack" "portfolio-api-lambda" {
  name          = "portfolio-api-lambda"
  template_body = file("template.yaml")
  capabilities  = ["CAPABILITY_IAM", "CAPABILITY_AUTO_EXPAND", "CAPABILITY_NAMED_IAM"]
}

data "aws_cloudformation_stack" "portfolio-api-lambda" {
  name = "portfolio-api-lambda"
  depends_on = [
    aws_cloudformation_stack.portfolio-api-lambda
  ]
}

resource "aws_ssm_parameter" "portfolio-dsn-production" {
  name        = "/portfolio/dsn"
  value       = "null"
  type        = "String"
  description = "DSN"
  lifecycle {
    ignore_changes = [value]
  }
}
resource "aws_ssm_parameter" "portfolio-dsn-qa" {
  name        = "/portfolio/dsn/qa"
  value       = "null"
  type        = "String"
  description = "DSN_QA"
  lifecycle {
    ignore_changes = [value]
  }
}

