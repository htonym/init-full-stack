
data "terraform_remote_state" "ops" {
  backend = "s3"
  config = {
    bucket  = var.state_bucket_ops
    key     = var.state_bucket_key_ops
    region  = var.aws_region
    profile = var.aws_profile
  }
}
