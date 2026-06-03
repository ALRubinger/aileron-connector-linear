+++
name = "team-membership-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-membership-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_membership_create"]

[match]
intent = "Creates a new team membership, adding a user to a team. Validates that the user is not already a member, the team is not archived or retired, and the requesting user has permission to add members."

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_membership_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

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

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier of the team associated with the membership."
required = true

[[inputs]]
name = "userId"
type = "String"
description = "The identifier of the user associated with the membership."
required = true
+++
