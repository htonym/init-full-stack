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
