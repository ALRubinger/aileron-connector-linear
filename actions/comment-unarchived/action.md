+++
name = "comment-unarchived"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/comment-unarchived@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["comment_unarchived"]

[match]
intent = "Triggered when a a comment is unarchived"

[[execute]]
id = "comment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "comment_unarchived"
idempotent = true
+++
