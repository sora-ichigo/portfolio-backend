provider "aws" {
  region = "ap-northeast-1"
}

provider "aws" {
  alias  = "virginia"
  region = "us-east-1"
}

terraform {
  backend "s3" {
    bucket = "tfstate-igsr5"
    key    = "igsr5-terraform/terraform.tfstate"
    region = "ap-northeast-1"
  }
}
