
resource "aws_cognito_user_pool" "this" {
  name = var.namespace

  alias_attributes = [
    "email",
    "preferred_username"
  ]

  password_policy {
    minimum_length                   = 8
    temporary_password_validity_days = 7
  }

  username_configuration {
    case_sensitive = false
  }

  mfa_configuration = "OFF"

  admin_create_user_config {
    allow_admin_create_user_only = true
  }

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }

  auto_verified_attributes = ["email"]

  user_attribute_update_settings {
    attributes_require_verification_before_update = ["email"]
  }

  schema {
    attribute_data_type      = "String"
    developer_only_attribute = false
    mutable                  = false
    name                     = "email"
    required                 = true
  }

  schema {
    attribute_data_type      = "String"
    developer_only_attribute = false
    mutable                  = true
    name                     = "family_name"
    required                 = true
  }

  schema {
    attribute_data_type      = "String"
    developer_only_attribute = false
    mutable                  = true
    name                     = "given_name"
    required                 = true
  }

  tags = {
    project = var.namespace
  }
}

resource "aws_cognito_user_pool_domain" "this" {
  user_pool_id = aws_cognito_user_pool.this.id
  domain       = var.namespace
}

# Staging App Client
resource "aws_cognito_user_pool_client" "staging_app" {
  user_pool_id    = aws_cognito_user_pool.this.id
  name            = "init-full-stack-app"
  generate_secret = true

  callback_urls = [
    "http://localhost:8000/api/callback",
    "https://${var.staging_app_domain}/api/callback",
  ]
  logout_urls = [
    "http://localhost:8000/",
    "https://${var.staging_app_domain}/",
  ]

  allowed_oauth_flows_user_pool_client = true
  supported_identity_providers         = ["COGNITO"]
  allowed_oauth_flows                  = ["code"]
  allowed_oauth_scopes                 = ["email", "openid"]
}
