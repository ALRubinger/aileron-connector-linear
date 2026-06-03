+++
name = "view-preferences-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/view-preferences-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["view_preferences_update"]

[match]
intent = "Updates an existing view preferences object. For user-type preferences, only the owning user can update them."

[[execute]]
id = "view"
connector = "github://ALRubinger/aileron-connector-linear"
op = "view_preferences_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the ViewPreferences object."
required = true

[[inputs]]
name = "insights"
type = "JSONObject"
description = "The default parameters for the insight on that view."
required = false

[[inputs]]
name = "preferences"
type = "JSONObject"
description = "View preferences."
required = false
+++
