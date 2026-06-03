+++
name = "organization-invite-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-invite-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_invite_delete"]

[match]
intent = "Deletes (archives) a workspace invite, preventing it from being accepted."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_invite_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the organization invite to delete."
required = true
+++
