
output "ec2_ami_public_ip" {
  value       = module.ec2_app.public_ip
  description = "Details of the private subnet"
}
