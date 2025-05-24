resource "aws_iam_instance_profile" "this" {
  name = "${var.namespace}-app-${var.environment}-profile"
  role = aws_iam_role.this.name
}

resource "aws_iam_role" "this" {
  name               = "${var.namespace}-app-${var.environment}"
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

data "aws_iam_policy_document" "extra_permissions" {
  statement {
    actions = [
      "ecr:GetAuthorizationToken",
      "ecr:BatchGetImage",
      "ecr:GetDownloadUrlForLayer",
      "ecr:DescribeImages",
      "ecr:DescribeRepositories",
      "ecr:ListImages",
      "ecr:ListTagsForResource",
      "ecr:DescribeImageScanFindings",
    ]

    resources = ["*"]
  }

  statement {
    actions = [
      "ssm:GetParameter",
      "ssm:GetParameters",
      "ssm:GetParameterHistory",
      "ssm:DescribeParameters"
    ]
    resources = [
      data.aws_ssm_parameter.config.arn
    ]
  }
}

resource "aws_iam_policy" "extra_permissions" {
  name        = "${var.namespace}-extra-permissions-${var.environment}"
  description = "Policy to allow extra permissions"
  policy      = data.aws_iam_policy_document.extra_permissions.json
}

resource "aws_iam_role_policy_attachment" "extra_permissions" {
  role       = aws_iam_role.this.name
  policy_arn = aws_iam_policy.extra_permissions.arn
}

data "aws_ssm_parameter" "config" {
  name            = "/${var.environment}/${var.namespace}/config"
  with_decryption = true
}
