+++
name = "favorite-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/favorite-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["favorite_create"]

[match]
intent = "Creates a new favorite for the authenticated user. Exactly one target entity must be specified. If a favorite for the same entity already exists, the existing favorite is returned (upsert behavior)."

[[execute]]
id = "favorite"
connector = "github://ALRubinger/aileron-connector-linear"
op = "favorite_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "customViewId"
type = "String"
description = "The identifier of the custom view to favorite."
required = false

[[inputs]]
name = "customerId"
type = "String"
description = "The identifier of the customer to favorite."
required = false

[[inputs]]
name = "cycleId"
type = "String"
description = "The identifier of the cycle to favorite."
required = false

[[inputs]]
name = "dashboardId"
type = "String"
description = "The identifier of the dashboard to favorite."
required = false

[[inputs]]
name = "documentId"
type = "String"
description = "The identifier of the document to favorite."
required = false

[[inputs]]
name = "facetId"
type = "String"
description = "The identifier of the facet to favorite."
required = false

[[inputs]]
name = "folderName"
type = "String"
description = "The name of the favorite folder."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "[INTERNAL] The identifier of the initiative to favorite."
required = false

[[inputs]]
name = "initiativeTab"
type = "InitiativeTab"
description = "The tab of the initiative to favorite."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The identifier of the issue to favorite. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = false

[[inputs]]
name = "labelId"
type = "String"
description = "The identifier of the label to favorite."
required = false

[[inputs]]
name = "parentId"
type = "String"
description = "The parent folder of the favorite."
required = false

[[inputs]]
name = "pipelineTab"
type = "PipelineTab"
description = "The tab of the release pipeline to favorite."
required = false

[[inputs]]
name = "predefinedViewTeamId"
type = "String"
description = "The identifier of team for the predefined view to favorite."
required = false

[[inputs]]
name = "predefinedViewType"
type = "String"
description = "The type of the predefined view to favorite."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "The identifier of the project to favorite."
required = false

[[inputs]]
name = "projectLabelId"
type = "String"
description = "The identifier of the project label to favorite."
required = false

[[inputs]]
name = "projectTab"
type = "ProjectTab"
description = "The tab of the project to favorite."
required = false

[[inputs]]
name = "pullRequestId"
type = "String"
description = "The identifier of the pull request to favorite."
required = false

[[inputs]]
name = "releaseId"
type = "String"
description = "The identifier of the release to favorite."
required = false

[[inputs]]
name = "releaseNoteId"
type = "String"
description = "The identifier of the release note to favorite."
required = false

[[inputs]]
name = "releasePipelineId"
type = "String"
description = "The identifier of the release pipeline to favorite."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The position of the item in the favorites list."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier of the team to favorite."
required = false

[[inputs]]
name = "userId"
type = "String"
description = "The identifier of the user to favorite."
required = false
+++
