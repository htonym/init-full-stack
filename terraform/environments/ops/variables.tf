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

variable "staging_app_domain" {
  type        = string
  description = "The domain for the staging app. Example: staging.widget.com"
}
