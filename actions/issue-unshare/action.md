+++
name = "issue-unshare"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-unshare@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_unshare"]

[match]
intent = "Stops sharing an issue with a user. The viewer must have native access to the issue's full sub-issue tree and permission to share issues in the issue's team. Issues that inherit sharing from a parent issue cannot be unshared directly."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_unshare"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to stop sharing."
required = true

[[inputs]]
name = "userId"
type = "String"
description = "The identifier of the user to remove from the issue share."
required = true
+++
