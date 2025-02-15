locals {
  namespace_env = "${var.namespace}-${var.environment}"
}

# EC2 instance
resource "aws_instance" "this" {
  ami           = var.ec2_ami # use AL2023 with Docker installed
  instance_type = "t2.micro"
  key_name      = var.ssh_key_pair

  subnet_id              = var.subnet_id
  vpc_security_group_ids = [aws_security_group.this.id]

  associate_public_ip_address = true

  root_block_device {
    delete_on_termination = true
    volume_size           = 10
    volume_type           = "gp3"
  }

  tags = {
    Name = "${var.namespace}-app-${var.environment}"
  }
}
