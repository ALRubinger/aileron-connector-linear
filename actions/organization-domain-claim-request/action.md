+++
name = "organization-domain-claim-request"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-domain-claim-request@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_domain_claim_request"]

[match]
intent = "[INTERNAL] Checks whether the domain can be claimed."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_domain_claim_request"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The ID of the organization domain to claim."
required = true
+++
