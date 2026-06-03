+++
name = "comment-resolve"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/comment-resolve@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["comment_resolve"]

[match]
intent = "Resolves a comment thread. Marks the root comment as resolved by the current user."

[[execute]]
id = "comment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "comment_resolve"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the comment to resolve."
required = true

[[inputs]]
name = "resolvingCommentId"
type = "String"
description = "The identifier of the child comment that resolves the thread. If not provided, the thread is resolved without a specific resolving comment."
required = false
+++
