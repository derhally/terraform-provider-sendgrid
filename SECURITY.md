# Security Policy

## Supported Versions

We release patches for security vulnerabilities for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 2.x     | :white_check_mark: |
| 1.x     | :x:                |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of the Terraform Provider for SendGrid seriously. If you believe you have found a security vulnerability, please report it to us as described below.

### Please Do Not

- **Do not** open a public GitHub issue for security vulnerabilities
- **Do not** publicly disclose the vulnerability until it has been addressed

### Please Do

**Report security vulnerabilities via GitHub Security Advisories:**

1. Go to the [Security tab](https://github.com/arslanbekov/terraform-provider-sendgrid/security/advisories)
2. Click "Report a vulnerability"
3. Fill out the form with as much detail as possible

**Alternatively, you can:**

- Email the maintainer directly through their GitHub profile: [@arslanbekov](https://github.com/arslanbekov)

### What to Include in Your Report

Please include the following information in your report:

- **Type of vulnerability** (e.g., credential exposure, injection, etc.)
- **Full paths** of source file(s) related to the manifestation of the issue
- **Location** of the affected source code (tag/branch/commit or direct URL)
- **Step-by-step instructions** to reproduce the issue
- **Proof-of-concept or exploit code** (if possible)
- **Impact** of the issue, including how an attacker might exploit it
- **Suggested fix** (if you have one)

### What to Expect

- **Acknowledgment**: We will acknowledge receipt of your vulnerability report within 48 hours
- **Communication**: We will keep you informed about our progress addressing the vulnerability
- **Timeline**: We aim to address critical vulnerabilities within 7-14 days
- **Credit**: We will credit you in the security advisory (unless you prefer to remain anonymous)

## Security Best Practices for Users

### API Key Management

**Never commit API keys to version control:**

```bash
# Use environment variables
export SENDGRID_API_KEY="SG.your-api-key"

# Or use terraform.tfvars (and add it to .gitignore)
echo "sendgrid_api_key = \"SG.your-api-key\"" > terraform.tfvars
echo "terraform.tfvars" >> .gitignore
```

**Use least privilege principle:**

```hcl
resource "sendgrid_api_key" "terraform" {
  name   = "terraform-automation"
  scopes = [
    "mail.send",
    "sender_verification_eligible"
  ]
}
```

**Rotate API keys regularly:**

```bash
# Create new key before revoking old one
# Update your secrets management system
# Test with new key
# Revoke old key only after confirming new one works
```

### Sensitive Data in State Files

Terraform state files may contain sensitive information:

```hcl
terraform {
  backend "s3" {
    bucket         = "my-terraform-state"
    key            = "sendgrid/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-locks"
  }
}
```

**Best practices:**

- Use remote state with encryption
- Enable state locking
- Restrict access to state files
- Use secret management tools (HashiCorp Vault, AWS Secrets Manager, etc.)

### CI/CD Security

**GitHub Actions example:**

```yaml
- name: Terraform Apply
  env:
    SENDGRID_API_KEY: ${{ secrets.SENDGRID_API_KEY }}
  run: terraform apply -auto-approve
```

**Best practices:**

- Store secrets in your CI/CD platform's secret management
- Never log sensitive values
- Use short-lived credentials when possible
- Audit access to secrets regularly

### Network Security

**Considerations:**

- Provider communicates with SendGrid API over HTTPS
- Ensure your network allows outbound HTTPS (443) traffic to SendGrid endpoints
- Consider using VPN or private networks for production deployments

### Dependency Security

**Keep dependencies updated:**

```bash
# Check for outdated dependencies
go list -u -m all

# Update dependencies
go get -u
go mod tidy
```

**Monitor security advisories:**

- Enable GitHub Dependabot alerts
- Subscribe to [Go security announcements](https://groups.google.com/g/golang-announce)
- Monitor [HashiCorp security bulletins](https://www.hashicorp.com/security)

## Known Security Considerations

### API Key Exposure

**Risk**: API keys stored in Terraform state files can be exposed if state is not properly secured.

**Mitigation**:

- Use remote state with encryption
- Implement proper access controls
- Consider using dynamic credentials or short-lived tokens
- Regularly rotate API keys

### Rate Limiting

**Risk**: Aggressive automation might trigger rate limiting or account suspension.

**Mitigation**:

- Provider implements exponential backoff
- Respect SendGrid's rate limits
- Use appropriate parallelism settings in Terraform

### Credential Leaks in Logs

**Risk**: Sensitive data might be logged during provider operations.

**Mitigation**:

- Provider marks sensitive attributes
- Review logs before sharing
- Use log sanitization in production environments

## Security Update Process

When a security vulnerability is confirmed:

1. **Assessment**: We assess the severity and impact
2. **Fix**: We develop a fix in a private branch
3. **Testing**: We thoroughly test the fix
4. **Release**: We release a patch version
5. **Disclosure**: We publish a security advisory with details
6. **Notification**: We notify affected users through:
   - GitHub Security Advisories
   - Release notes
   - GitHub Discussions (for major issues)

## Scope

This security policy applies to:

- The Terraform Provider for SendGrid codebase
- Official releases and distributions
- Documentation and examples

This policy does not cover:

- SendGrid's API or infrastructure (report to SendGrid directly)
- Terraform core (report to HashiCorp)
- Third-party tools or integrations

## Additional Resources

- [SendGrid Security](https://sendgrid.com/en-us/policies/security)
- [Terraform Security](https://www.hashicorp.com/en/trust/security)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [CWE Top 25](https://cwe.mitre.org/top25/)

## Questions?

If you have questions about this security policy, please open a [Discussion](https://github.com/arslanbekov/terraform-provider-sendgrid/discussions) or contact the maintainers.

---

Thank you for helping keep the Terraform Provider for SendGrid and its users safe!
