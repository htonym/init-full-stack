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
