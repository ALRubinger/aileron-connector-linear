+++
name = "view-preferences-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/view-preferences-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["view_preferences_create"]

[match]
intent = "Creates a new view preferences object. If conflicting preferences already exist for the same view type and scope, the existing preferences are replaced."

[[execute]]
id = "view"
connector = "github://ALRubinger/aileron-connector-linear"
op = "view_preferences_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "customViewId"
type = "String"
description = "The custom view these view preferences are associated with."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "[Internal] The initiative these view preferences are associated with."
required = false

[[inputs]]
name = "insights"
type = "JSONObject"
description = "The default parameters for the insight on that view."
required = false

[[inputs]]
name = "labelId"
type = "String"
description = "The label these view preferences are associated with."
required = false

[[inputs]]
name = "preferences"
type = "JSONObject"
description = "View preferences object."
required = true

[[inputs]]
name = "projectId"
type = "String"
description = "The project these view preferences are associated with."
required = false

[[inputs]]
name = "projectLabelId"
type = "String"
description = "The project label these view preferences are associated with."
required = false

[[inputs]]
name = "releasePipelineId"
type = "String"
description = "The release pipeline these view preferences are associated with."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The team these view preferences are associated with."
required = false

[[inputs]]
name = "type"
type = "ViewPreferencesType"
description = "The type of view preferences (either user or workspace level preferences)."
required = true

[[inputs]]
name = "userId"
type = "String"
description = "The user profile these view preferences are associated with."
required = false

[[inputs]]
name = "viewType"
type = "ViewType"
description = "The view type of the view preferences are associated with."
required = true
+++
