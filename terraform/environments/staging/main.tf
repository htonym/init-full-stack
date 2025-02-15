provider "aws" {
  region  = var.aws_region
  profile = var.aws_profile
}

terraform {
  backend "s3" {
    profile = "management-admin"
    # bucket and region defined in backend.conf
  }
}

locals {
  environment   = "staging"
  namespace_env = "${var.namespace}-${local.environment}"
  ops_state     = data.terraform_remote_state.ops.outputs
}

module "ec2_app" {
  source = "../../modules/ec2_app"

  subnet_id      = local.ops_state.public_subnet.id
  namespace      = var.namespace
  aws_region     = var.aws_region
  environment    = local.environment
  vpc_id         = local.ops_state.vpc_id
  allowed_ssh_ip = var.allowed_ssh_ip
  ssh_key_pair   = var.ssh_key_pair
  ec2_ami        = var.ec2_ami
}
