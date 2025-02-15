variable "aws_region" {
  type = string
}

variable "environment" {
  type = string
}

variable "namespace" {
  type        = string
  description = "Name for the infrastructure ops"
}

variable "subnet_id" {
  type = string
}

variable "allowed_ssh_ip" {
  type        = string
  description = "The public IP of the machine that will login into EC2 via ssh"
}

variable "vpc_id" {
  type        = string
  description = "Provide VPC ID to be used for the EC2 instance."
}

variable "ssh_key_pair" {
  type = string
}

variable "ec2_ami" {
  type        = string
  description = "Amazon Machine Image with docker (AL2023)"
}
