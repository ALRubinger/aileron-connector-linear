+++
name = "custom-view-has-subscribers"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/custom-view-has-subscribers@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["custom_view_has_subscribers"]

[match]
intent = "Whether a custom view has active notification subscribers other than the current user. Useful for warning before deleting a shared view."

[[execute]]
id = "custom"
connector = "github://ALRubinger/aileron-connector-linear"
op = "custom_view_has_subscribers"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the custom view."
required = true
+++
