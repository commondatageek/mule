terraform {
  required_version = "1.1.4"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# for easy local development
provider "aws" {
  region  = "us-west-2"
  profile = "dev"
}