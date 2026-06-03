+++
name = "user-flag-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-flag-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_flag_update"]

[match]
intent = "Updates a specific user settings flag by performing an operation (e.g., increment or clear) on it."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_flag_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "flag"
type = "UserFlagType"
description = "Settings flag to increment."
required = true

[[inputs]]
name = "operation"
type = "UserFlagUpdateOperation"
description = "Flag operation to perform."
required = true
+++
