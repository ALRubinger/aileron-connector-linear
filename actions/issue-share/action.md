+++
name = "issue-share"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-share@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_share"]

[match]
intent = "Shares an issue with a user who would not otherwise be able to access it. The viewer must have native access to the issue's full sub-issue tree and permission to share issues in the issue's team. Issues that inherit sharing from a parent issue cannot be shared directly."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_share"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to share."
required = true

[[inputs]]
name = "userId"
type = "String"
description = "The identifier of the user to share the issue with."
required = true
+++
