+++
name = "custom-view-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/custom-view-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["custom_view_delete"]

[match]
intent = "Deletes a custom view."

[[execute]]
id = "custom"
connector = "github://ALRubinger/aileron-connector-linear"
op = "custom_view_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the custom view to delete."
required = true
+++
