+++
name = "organization-invite-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-invite-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_invite_create"]

[match]
intent = "Creates a new workspace invite and sends an invitation email to the specified address. The invite includes a role assignment and optional team memberships."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_invite_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "email"
type = "String"
description = "The email of the invitee."
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "metadata"
type = "JSONObject"
description = "[INTERNAL] Optional metadata about the invite."
required = false

[[inputs]]
name = "role"
type = "UserRoleType"
description = "What user role the invite should grant."
required = false

[[inputs]]
name = "teamIds"
type = "String"
description = "The teams that the user has been invited to."
required = false
+++
