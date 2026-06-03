+++
name = "issue-subscribe"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-subscribe@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_subscribe"]

[match]
intent = "Subscribes a user to an issue."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_subscribe"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to subscribe to."
required = true

[[inputs]]
name = "userEmail"
type = "String"
description = "The email of the user to subscribe, default is the current user."
required = false

[[inputs]]
name = "userId"
type = "String"
description = "The identifier of the user to subscribe, default is the current user."
required = false
+++
