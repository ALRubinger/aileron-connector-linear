+++
name = "organization-domain-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-domain-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_domain_delete"]

[match]
intent = "Deletes a domain."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_domain_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the domain to delete."
required = true
+++
