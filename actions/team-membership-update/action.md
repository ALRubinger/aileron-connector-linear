+++
name = "team-membership-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-membership-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_membership_update"]

[match]
intent = "Updates a team membership, such as changing ownership status or sort order."

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_membership_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the team membership to update."
required = true

[[inputs]]
name = "owner"
type = "Boolean"
description = "Internal. Whether the user is the owner of the team."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The position of the item in the users list."
required = false
+++
