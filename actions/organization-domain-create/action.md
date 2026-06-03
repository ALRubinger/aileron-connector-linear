+++
name = "organization-domain-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-domain-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_domain_create"]

[match]
intent = "[INTERNAL] Adds a domain to be allowed for a workspace."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_domain_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "authType"
type = "String"
description = "The authentication type this domain is for."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "identityProviderId"
type = "String"
description = "The identity provider to which to add the domain."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The domain name to add."
required = true

[[inputs]]
name = "triggerEmailVerification"
type = "Boolean"
description = "Whether to trigger an email verification flow during domain creation."
required = false

[[inputs]]
name = "verificationEmail"
type = "String"
description = "The email address to which to send the verification code."
required = false
+++
