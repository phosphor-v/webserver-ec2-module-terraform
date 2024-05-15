provider "aws" {
  region = "eu-north-1"
}

terraform {
  backend "s3" {
    bucket         = "phosphor-tfstate"
    key            = "week10/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
  }
}