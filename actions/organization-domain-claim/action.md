+++
name = "organization-domain-claim"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-domain-claim@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_domain_claim"]

[match]
intent = "[INTERNAL] Verifies a domain claim."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_domain_claim"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The ID of the organization domain to claim."
required = true
+++
