
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

  vpc_name                    = var.ops_name
  aws_region                  = var.aws_region
  vpc_cidr_block              = "10.0.0.0/16"
  public_subnet_cidr_block    = "10.0.1.0/24"
  private_subnet_cidr_block_a = "10.0.2.0/24"
  private_subnet_cidr_block_b = "10.0.3.0/24"
}
