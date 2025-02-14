
variable "namespace" {
  description = "Name of the VPC"
  type        = string
}

variable "aws_region" {
  description = "AWS Region"
  type        = string
}

variable "environment" {
  description = "Infrastructure Environment"
  type        = string
}

variable "vpc_cidr_block" {
  description = "value"
  type        = string
}

variable "public_subnet_cidr_block" {
  description = "value"
  type        = string
}

variable "private_subnet_cidr_block_a" {
  description = "value"
  type        = string
}

variable "private_subnet_cidr_block_b" {
  description = "value"
  type        = string
}
