+++
name = "organization-domain-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-domain-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_domain_update"]

[match]
intent = "[INTERNAL] Updates a workspace domain's settings."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_domain_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "disableOrganizationCreation"
type = "Boolean"
description = "Prevent users with this domain to create new workspaces. Only allowed to set on claimed domains!"
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the domain to update."
required = true
+++
