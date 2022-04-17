resource "aws_s3_bucket" "igsr5-portfolio-api-lambda-code" {
  bucket = "igsr5-portfolio-api-lambda-code"
}

resource "aws_cloudformation_stack" "igsr5-portfolio-api-lambda" {
  name          = "igsr5-portfolio-api-lambda"
  template_body = file("template.yaml")
  capabilities  = ["CAPABILITY_IAM", "CAPABILITY_AUTO_EXPAND"]
}

data "aws_cloudformation_stack" "igsr5-portfolio-api-lambda" {
  name = "igsr5-portfolio-api-lambda"
  depends_on = [
    aws_cloudformation_stack.igsr5-portfolio-api-lambda
  ]
}
