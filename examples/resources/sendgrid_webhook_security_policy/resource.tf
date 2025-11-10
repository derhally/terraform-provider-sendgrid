# OAuth (OIDC) only security policy
resource "sendgrid_webhook_security_policy" "oauth_only" {
  name = "OAuth Policy"

  oauth {
    client_id     = "my-client-id"
    client_secret = var.oauth_client_secret
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

# Combined OAuth and Signature security policy
resource "sendgrid_webhook_security_policy" "combined" {
  name = "Combined Security Policy"

  oauth {
    client_id     = "my-client-id"
    client_secret = var.oauth_client_secret
    token_url     = "https://oauth.example.com/token"
    scopes        = ["webhooks:read", "webhooks:write"]
  }

  signature {
    enabled = true
  }
}

# Output the public key for signature verification
output "signature_public_key" {
  description = "Public key for webhook signature verification"
  value       = sendgrid_webhook_security_policy.signature_only.signature[0].public_key
  sensitive   = true
}

output "combined_public_key" {
  description = "Public key for combined policy signature verification"
  value       = sendgrid_webhook_security_policy.combined.signature[0].public_key
  sensitive   = true
}

