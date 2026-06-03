+++
name = "comment-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/comment-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["comment_created"]

[match]
intent = "Triggered when a comment is created"

[[execute]]
id = "comment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "comment_created"
idempotent = true
+++
