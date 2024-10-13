terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.71.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "2.4.2"
    }
    null = {
      source  = "hashicorp/null"
      version = "3.2.2"
    }
  }
}

provider "aws" {
  region = "eu-central-1"
}
