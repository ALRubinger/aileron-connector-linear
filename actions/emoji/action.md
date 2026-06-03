+++
name = "emoji"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/emoji@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["emoji"]

[match]
intent = "A specific custom emoji by ID or name."

[[execute]]
id = "emoji"
connector = "github://ALRubinger/aileron-connector-linear"
op = "emoji"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier or the name of the emoji to retrieve."
required = true
+++
