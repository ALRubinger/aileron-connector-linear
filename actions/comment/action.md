+++
name = "comment"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/comment@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["comment"]

[match]
intent = "A specific comment."

[[execute]]
id = "comment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "comment"
idempotent = true

[[inputs]]
name = "hash"
type = "String"
description = "The hash of the comment to retrieve."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the comment to retrieve."
required = false
+++
