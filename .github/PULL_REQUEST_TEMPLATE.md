## Description

<!-- Provide a brief description of your changes -->

## Type of Change

<!-- Mark the relevant option with an "x" -->

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Code refactoring (no functional changes)
- [ ] Test improvements
- [ ] Chore (dependency updates, tooling, etc.)

## Related Issues

<!-- Link to related issues using keywords: Fixes #123, Closes #456, Relates to #789 -->

Fixes #

## Changes Made

<!-- Describe the changes in detail -->

-
-
-

## Testing

<!-- Describe how you tested your changes -->

### Unit Tests

- [ ] I have added unit tests for my changes
- [ ] All unit tests pass locally (`make test`)

### Acceptance Tests

- [ ] I have added/updated acceptance tests (`make testacc`)
- [ ] All acceptance tests pass with a real SendGrid account
- [ ] Tests include Create, Read, Update, and Delete operations (if applicable)

### Manual Testing

<!-- Describe any manual testing performed -->

```hcl
# Example configuration used for testing
resource "sendgrid_resource" "test" {
  # ...
}
```

**Test Results:**

```
# Paste test output here
```

## Documentation

<!-- Check all that apply -->

- [ ] I have updated the documentation in `docs/`
- [ ] I have added examples to `examples/resources/`
- [ ] I have updated the CHANGELOG.md
- [ ] I have added/updated code comments where necessary
- [ ] All exported functions have godoc comments

## Checklist

<!-- Ensure all items are completed before requesting review -->

- [ ] My code follows the project's [coding standards](CONTRIBUTING.md#coding-standards)
- [ ] I have run `make fmt` to format the code
- [ ] I have run `make lint` and fixed all issues
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] My changes generate no new warnings or errors
- [ ] I have checked for any breaking changes
- [ ] My commit messages follow the [Conventional Commits](https://www.conventionalcommits.org/) format

## Breaking Changes

<!-- If this PR introduces breaking changes, describe them here and provide migration instructions -->

**Is this a breaking change?** <!-- Yes/No -->

<!-- If yes, describe the breaking changes and how users should migrate -->

## Screenshots (if applicable)

<!-- Add screenshots to help explain your changes -->

## Additional Context

<!-- Add any other context about the PR here -->

## Reviewer Notes

<!-- Any specific areas you'd like reviewers to focus on? -->

---

By submitting this pull request, I confirm that my contribution is made under the terms of the Mozilla Public License 2.0.
