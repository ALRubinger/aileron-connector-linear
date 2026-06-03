+++
name = "emoji-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/emoji-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["emoji_create"]

[match]
intent = "Creates a custom emoji."

[[execute]]
id = "emoji"
connector = "github://ALRubinger/aileron-connector-linear"
op = "emoji_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the custom emoji."
required = true

[[inputs]]
name = "url"
type = "String"
description = "The URL for the emoji."
required = true
+++
