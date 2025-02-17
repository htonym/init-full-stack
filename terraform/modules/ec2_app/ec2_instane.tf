locals {
  namespace_env = "${var.namespace}-${var.environment}"
}

# EC2 instance
resource "aws_instance" "this" {
  ami           = var.ec2_ami # use AL2023 with Docker installed
  instance_type = var.instance_type
  key_name      = var.ssh_key_pair

  subnet_id              = var.subnet_id
  vpc_security_group_ids = [aws_security_group.this.id]

  associate_public_ip_address = true

  iam_instance_profile = aws_iam_instance_profile.this.name

  root_block_device {
    delete_on_termination = true
    volume_size           = 10
    volume_type           = "gp3"
  }

  user_data = base64encode(templatefile("${path.module}/user_data.sh.tpl", {
    sub_domain      = var.sub_domain
    caddy_ecr_image = var.caddy_ecr_image
    aws_region      = var.aws_region
    aws_account_id  = var.aws_account_id
  }))

  tags = {
    Name = "${var.namespace}-app-${var.environment}"
  }
}

resource "null_resource" "wait_for_ec2_running" {
  depends_on = [aws_instance.this]
  provisioner "local-exec" {
    command = "aws --profile ${var.aws_profile} ec2 wait instance-running --instance-ids ${aws_instance.this.id}"
  }
}
