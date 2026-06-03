+++
name = "user-unsuspend"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-unsuspend@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_unsuspend"]

[match]
intent = "Re-activates a suspended user, restoring their access to the workspace. Can only be called by a workspace admin or owner."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_unsuspend"
idempotent = false

[approval]
required = true

[[inputs]]
name = "forceBypassScimRestrictions"
type = "Boolean"
description = "[INTERNAL] Whether to bypass SCIM restrictions when unsuspending. Use with caution — this overrides identity provider management."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the user to unsuspend."
required = true
+++
