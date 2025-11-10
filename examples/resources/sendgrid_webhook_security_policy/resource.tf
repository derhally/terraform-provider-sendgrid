# OAuth (OIDC) only security policy
resource "sendgrid_webhook_security_policy" "oauth_only" {
  name = "OAuth Policy"

  oauth {
    client_id     = "my-client-id"
    client_secret = "my-client-secret"
    token_url     = "https://oauth.example.com/token"
    scopes        = ["webhooks:read", "webhooks:write"]
  }
}

# Signature verification only security policy
resource "sendgrid_webhook_security_policy" "signature_only" {
  name = "Signature Verification Policy"

  signature {
    enabled = true
  }
}

# The public key can be referenced for verification
output "public_key" {
  value     = sendgrid_webhook_security_policy.signature_only.signature[0].public_key
  sensitive = true
}

# Combined OAuth and Signature security policy
resource "sendgrid_webhook_security_policy" "both" {
  name = "Combined Security Policy"

  oauth {
    client_id     = "my-client-id"
    client_secret = "my-client-secret"
    token_url     = "https://oauth.example.com/token"
    scopes        = ["webhooks:read", "webhooks:write"]
  }

  signature {
    enabled = true
  }
}

# Attached to Parse Webhook
resource "sendgrid_webhook_security_policy" "parse_security" {
  name = "Parse Webhook Security"

  signature {
    enabled = true
  }
}

resource "sendgrid_parse_webhook" "example" {
  hostname                   = "parse.example.com"
  url                        = "https://api.example.com/parse"
  spam_check                 = true
  send_raw                   = false
  webhook_security_policy_id = sendgrid_webhook_security_policy.parse_security.id
}
