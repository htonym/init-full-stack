resource "aws_security_group" "this" {
  name_prefix = local.namespace_env
  vpc_id      = var.vpc_id

  tags = {
    Name = local.namespace_env
  }
}

resource "aws_security_group_rule" "outbound_https" {
  security_group_id = aws_security_group.this.id
  description       = "Allow to EC2 to make HTTPS request to other services such as AWS services"

  type        = "egress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

// Allow inbound http for testing 
resource "aws_security_group_rule" "inbound_http" {
  security_group_id = aws_security_group.this.id
  description       = "Allow inbound http for testing "

  type        = "ingress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "allow_ssh" {
  security_group_id = aws_security_group.this.id

  type        = "ingress"
  from_port   = 22
  to_port     = 22
  protocol    = "tcp"
  cidr_blocks = ["${var.allowed_ssh_ip}/32"]

  description = "Allow SSH access"
}

resource "aws_security_group_rule" "allow_icmp" {
  security_group_id = aws_security_group.this.id
  description       = "Allow ICMP traffic for ping"

  type        = "ingress"
  from_port   = -1
  to_port     = -1
  protocol    = "icmp"
  cidr_blocks = ["0.0.0.0/0"]
}
