+++
name = "comment-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/comment-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["comment_update"]

[match]
intent = "Updates a comment."

[[execute]]
id = "comment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "comment_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "body"
type = "String"
description = "The comment content."
required = false

[[inputs]]
name = "bodyData"
type = "JSON"
description = "[Internal] The comment content as a Prosemirror document."
required = false

[[inputs]]
name = "doNotSubscribeToIssue"
type = "Boolean"
description = "[INTERNAL] Flag to prevent auto subscription to the issue the comment is updated on."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the comment to update."
required = true

[[inputs]]
name = "quotedText"
type = "String"
description = "The text that this comment references. Only defined for inline comments."
required = false

[[inputs]]
name = "resolvingCommentId"
type = "String"
description = "[INTERNAL] The child comment that resolves this thread."
required = false

[[inputs]]
name = "resolvingUserId"
type = "String"
description = "[INTERNAL] The user who resolved this thread."
required = false

[[inputs]]
name = "skipEditedAt"
type = "Boolean"
description = "[INTERNAL] Flag to prevent setting editedAt when updating bodyData (used for background uploads)."
required = false

[[inputs]]
name = "subscriberIds"
type = "String"
description = "[INTERNAL] The identifiers of the users subscribing to this comment."
required = false
+++
