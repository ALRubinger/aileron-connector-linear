+++
name = "comment-unresolve"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/comment-unresolve@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["comment_unresolve"]

[match]
intent = "Unresolves a previously resolved comment thread. Clears the resolved state on the root comment."

[[execute]]
id = "comment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "comment_unresolve"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the comment to unresolve."
required = true
+++
