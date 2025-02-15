output "private_subnet_a" {
  value       = module.network.private_subnet_a
  description = "Details of the private subnet"
}

output "private_subnet_b" {
  value       = module.network.private_subnet_b
  description = "Details of the private subnet"
}

output "public_subnet" {
  value       = module.network.public_subnet
  description = "Details of the private subnet"
}

output "vpc_id" {
  value = module.network.vpc_id
}
