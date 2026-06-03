+++
name = "organization-domain-verify"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-domain-verify@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_domain_verify"]

[match]
intent = "[INTERNAL] Verifies a domain to be added to a workspace."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_domain_verify"
idempotent = false

[approval]
required = true

[[inputs]]
name = "organizationDomainId"
type = "String"
description = "The identifier in UUID v4 format of the domain being verified."
required = true

[[inputs]]
name = "verificationCode"
type = "String"
description = "The verification code sent via email."
required = true
+++
