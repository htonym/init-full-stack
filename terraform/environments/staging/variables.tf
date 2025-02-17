variable "aws_region" {
  type = string
}

variable "aws_profile" {
  type = string
}

variable "namespace" {
  type        = string
  description = "Name for the infrastructure ops"
}

variable "state_bucket_ops" {
  type        = string
  description = "name of s3 Ops state bucket"
}

variable "state_bucket_key_ops" {
  type        = string
  description = "name of s3 Ops state bucket key"
}

variable "allowed_ssh_ip" {
  type        = string
  description = "The public IP of the machine that will login into EC2 via ssh"
}

variable "ssh_key_pair" {
  type        = string
  description = "The public IP of the machine that will login into EC2 via ssh"
}

variable "ec2_ami" {
  type        = string
  description = "Amazon Machine Image with docker (AL2023)"
}

variable "app_domain" {
  type        = string
  description = "existing route 53 domain name"
}

variable "caddy_ecr_image" {
  type        = string
  description = "Example: your-account-id.dkr.ecr.your-region.amazonaws.com/your-repo-name:latest"
}

variable "aws_account_id" {
  type = string
}

variable "app_ecr_image" {
  type = string
}
