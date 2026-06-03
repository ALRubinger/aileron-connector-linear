+++
name = "template"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/template@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["template"]

[match]
intent = "A specific template."

[[execute]]
id = "template"
connector = "github://ALRubinger/aileron-connector-linear"
op = "template"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the template to retrieve."
required = true
+++
