+++
name = "custom-view-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/custom-view-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["custom_view_update"]

[match]
intent = "Updates a custom view."

[[execute]]
id = "custom"
connector = "github://ALRubinger/aileron-connector-linear"
op = "custom_view_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The color of the icon of the custom view."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the custom view."
required = false

[[inputs]]
name = "feedItemFilterData"
type = "FeedItemFilter"
description = "The feed item filter applied to issues in the custom view."
required = false

[[inputs]]
name = "filterData"
type = "IssueFilter"
description = "The filter applied to issues in the custom view."
required = false

[[inputs]]
name = "icon"
type = "String"
description = "The icon of the custom view."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the custom view to update."
required = true

[[inputs]]
name = "initiativeFilterData"
type = "InitiativeFilter"
description = "[ALPHA] The initiative filter applied to issues in the custom view."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "[Internal] The id of the initiative associated with the custom view."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the custom view."
required = false

[[inputs]]
name = "ownerId"
type = "String"
description = "The owner of the custom view."
required = false

[[inputs]]
name = "projectFilterData"
type = "ProjectFilter"
description = "The project filter applied to issues in the custom view."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "[Internal] The id of the project associated with the custom view."
required = false

[[inputs]]
name = "shared"
type = "Boolean"
description = "Whether the custom view is shared with everyone in the workspace."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The id of the team associated with the custom view."
required = false
+++
