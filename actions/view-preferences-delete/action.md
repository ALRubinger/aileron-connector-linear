+++
name = "view-preferences-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/view-preferences-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["view_preferences_delete"]

[match]
intent = "Deletes a view preferences object. If the preferences do not exist, the operation is treated as a successful idempotent deletion."

[[execute]]
id = "view"
connector = "github://ALRubinger/aileron-connector-linear"
op = "view_preferences_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the ViewPreferences to delete."
required = true
+++
