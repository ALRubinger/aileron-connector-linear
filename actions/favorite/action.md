+++
name = "favorite"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/favorite@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["favorite"]

[match]
intent = "A specific favorite by ID."

[[execute]]
id = "favorite"
connector = "github://ALRubinger/aileron-connector-linear"
op = "favorite"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
