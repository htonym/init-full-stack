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
  sub_domain    = "staging.${data.aws_route53_zone.app_domain.name}"
}

module "ec2_app" {
  source = "../../modules/ec2_app"

  app_version = "0.3.0-dev.4"

  port            = "8000"
  subnet_id       = local.ops_state.public_subnet.id
  namespace       = var.namespace
  aws_region      = var.aws_region
  environment     = local.environment
  vpc_id          = local.ops_state.vpc_id
  allowed_ssh_ip  = var.allowed_ssh_ip
  ssh_key_pair    = var.ssh_key_pair
  ec2_ami         = var.ec2_ami
  aws_profile     = var.aws_profile
  sub_domain      = local.sub_domain
  instance_type   = "t3.micro"
  aws_account_id  = var.aws_account_id
  caddy_ecr_image = var.caddy_ecr_image
  app_ecr_repo    = var.app_ecr_repo
}

resource "aws_route53_record" "staging_a_record" {
  zone_id = data.aws_route53_zone.app_domain.zone_id
  name    = local.sub_domain
  type    = "A"
  ttl     = 300
  records = [module.ec2_app.public_ip]

  # Avoid needing elastic ip by waiting for ec2 ip. Must always create and
  # destroy route and ec2 together.
  depends_on = [module.ec2_app.wait_for_ec2_running]
}

