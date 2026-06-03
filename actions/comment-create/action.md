+++
name = "comment-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/comment-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["create_comment"]

[match]
intent = "Add a comment to an existing Linear issue. Requires per-call user approval."

[[execute]]
id = "comment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "comment_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "body"
type = "String"
description = "Markdown body of the comment."
required = true

[[inputs]]
name = "issueId"
type = "String"
description = "UUID of the issue to comment on."
required = true
+++
