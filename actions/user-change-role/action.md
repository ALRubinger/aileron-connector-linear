+++
name = "user-change-role"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-change-role@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_change_role"]

[match]
intent = "Changes the workspace role of a user. The requesting user must have a role equal to or higher than the target role. Requires workspace admin permissions."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_change_role"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the user"
required = true

[[inputs]]
name = "role"
type = "UserRoleType"
description = "The new role for the user."
required = true
+++
