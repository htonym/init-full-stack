
# The Ops environment is intended to create resources that can be reused by
# other environments like: ECR repo or user pool.

provider "aws" {
  region  = var.aws_region
  profile = var.aws_profile
}

terraform {
  backend "s3" {
    key     = "init-full-stack-ops.tfstate"
    profile = "management-admin"
    # bucket and region defined in backend.conf
  }
}

locals {
  environment = "ops"
}


module "network" {
  source = "../../modules/network"

  namespace                   = var.namespace
  aws_region                  = var.aws_region
  environment                 = local.environment
  vpc_cidr_block              = "10.0.0.0/16"
  public_subnet_cidr_block    = "10.0.1.0/24"
  private_subnet_cidr_block_a = "10.0.2.0/24"
  private_subnet_cidr_block_b = "10.0.3.0/24"
}

resource "aws_ecr_repository" "app" {
  name = var.namespace

  image_scanning_configuration {
    scan_on_push = false
  }
}

resource "aws_ecr_repository" "caddy" {
  name                 = "caddy"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = false
  }
}
