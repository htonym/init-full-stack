
output "public_ip" {
  value = aws_instance.this.public_ip
}

output "wait_for_ec2_running" {
  value = null_resource.wait_for_ec2_running
}
