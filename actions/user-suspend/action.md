+++
name = "user-suspend"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-suspend@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_suspend"]

[match]
intent = "Suspends a user, deactivating their account and revoking access to the workspace. The suspended user's sessions are invalidated. Can only be called by a workspace admin or owner."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_suspend"
idempotent = false

[approval]
required = true

[[inputs]]
name = "forceBypassScimRestrictions"
type = "Boolean"
description = "[INTERNAL] Whether to bypass SCIM restrictions when suspending. Use with caution — this overrides identity provider management."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the user to suspend."
required = true
+++
