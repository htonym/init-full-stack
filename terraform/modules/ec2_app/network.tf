resource "aws_security_group" "service" {
  name_prefix = local.namespace_env
  vpc_id      = data.terraform_remote_state.ops.outputs.vpc_id

  tags = {
    Name = local.namespace_env
  }
}

resource "aws_security_group_rule" "outbound_https" {
  security_group_id = aws_security_group.service.id
  description       = "Allow to EC2 to make HTTPS request to other services such as AWS services"

  type        = "egress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

// Allow inbound http for testing 
resource "aws_security_group_rule" "inbound_http" {
  security_group_id = aws_security_group.service.id
  description       = "Allow inbound http for testing "

  type        = "ingress"
  from_port   = 80
  to_port     = 80
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

