+++
name = "organization-invite-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-invite-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_invite_update"]

[match]
intent = "Updates an existing workspace invite, such as changing the teams the invitee will be added to."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_invite_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the organization invite to update."
required = true

[[inputs]]
name = "teamIds"
type = "String"
description = "The teams that the user has been invited to."
required = true
+++
