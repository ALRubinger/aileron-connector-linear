+++
name = "team-membership-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-membership-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_membership_delete"]

[match]
intent = "Deletes a team membership, removing the user from the team. Users can remove their own membership, or team owners and workspace admins can remove other members."

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_membership_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "alsoLeaveParentTeams"
type = "Boolean"
description = "Whether to leave the parent teams."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the team membership to delete."
required = true
+++
