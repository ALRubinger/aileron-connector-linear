+++
name = "emoji-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/emoji-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["emoji_delete"]

[match]
intent = "Deletes an emoji."

[[execute]]
id = "emoji"
connector = "github://ALRubinger/aileron-connector-linear"
op = "emoji_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the emoji to delete."
required = true
+++
