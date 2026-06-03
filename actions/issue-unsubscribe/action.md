+++
name = "issue-unsubscribe"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-unsubscribe@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_unsubscribe"]

[match]
intent = "Unsubscribes a user from an issue."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_unsubscribe"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to unsubscribe from."
required = true

[[inputs]]
name = "userEmail"
type = "String"
description = "The email of the user to unsubscribe, default is the current user."
required = false

[[inputs]]
name = "userId"
type = "String"
description = "The identifier of the user to unsubscribe, default is the current user."
required = false
+++
