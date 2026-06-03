+++
name = "comment-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/comment-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["comment_delete"]

[match]
intent = "Deletes a comment."

[[execute]]
id = "comment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "comment_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the comment to delete."
required = true
+++
