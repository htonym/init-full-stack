resource "aws_iam_role" "this" {
  name               = "${var.environment}-westplexinvent-ec2-host-role"
  assume_role_policy = data.aws_iam_policy_document.ec2_access_role.json
}

data "aws_iam_policy_document" "ec2_access_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }
  }
}
