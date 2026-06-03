+++
name = "template-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/template-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["template_delete"]

[match]
intent = "Deletes a template."

[[execute]]
id = "template"
connector = "github://ALRubinger/aileron-connector-linear"
op = "template_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the template to delete."
required = true
+++
