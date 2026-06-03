+++
name = "custom-view"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/custom-view@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["custom_view"]

[match]
intent = "One specific custom view, looked up by ID or slug."

[[execute]]
id = "custom"
connector = "github://ALRubinger/aileron-connector-linear"
op = "custom_view"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the custom view to retrieve."
required = true
+++
