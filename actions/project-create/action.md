+++
name = "project-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_create"]

[match]
intent = "Creates a new project."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The color of the project."
required = false

[[inputs]]
name = "content"
type = "String"
description = "The project content as markdown."
required = false

[[inputs]]
name = "convertedFromIssueId"
type = "String"
description = "The ID of the issue that was converted into this project."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description for the project."
required = false

[[inputs]]
name = "icon"
type = "String"
description = "The icon of the project."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "labelIds"
type = "String"
description = "The identifiers of the project labels associated with this project."
required = false

[[inputs]]
name = "lastAppliedTemplateId"
type = "String"
description = "The ID of the last template applied to the project."
required = false

[[inputs]]
name = "leadId"
type = "String"
description = "The identifier of the project lead."
required = false

[[inputs]]
name = "memberIds"
type = "String"
description = "The identifiers of the members of this project."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the project."
required = true

[[inputs]]
name = "priority"
type = "Int"
description = "The priority of the project. 0 = No priority, 1 = Urgent, 2 = High, 3 = Medium, 4 = Low."
required = false

[[inputs]]
name = "prioritySortOrder"
type = "Float"
description = "The sort order for the project within shared views, when ordered by priority."
required = false

[[inputs]]
name = "slackChannelName"
type = "String"
description = "The full name for the Slack channel to create (including prefix). When provided, a Slack channel will be created and connected to the project."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order for the project within shared views."
required = false

[[inputs]]
name = "startDate"
type = "TimelessDate"
description = "The planned start date of the project."
required = false

[[inputs]]
name = "startDateResolution"
type = "DateResolutionType"
description = "The resolution of the project's start date."
required = false

[[inputs]]
name = "statusId"
type = "String"
description = "The ID of the project status."
required = false

[[inputs]]
name = "targetDate"
type = "TimelessDate"
description = "The planned target date of the project."
required = false

[[inputs]]
name = "targetDateResolution"
type = "DateResolutionType"
description = "The resolution of the project's estimated completion date."
required = false

[[inputs]]
name = "teamIds"
type = "String"
description = "The identifiers of the teams this project is associated with."
required = true

[[inputs]]
name = "templateId"
type = "String"
description = "The ID of a project template to apply when creating the project. Overrides useDefaultTemplate if both are provided."
required = false

[[inputs]]
name = "useDefaultTemplate"
type = "Boolean"
description = "When set to true, the default project template of the first team provided will be applied. If templateId is provided, this will be ignored."
required = false
+++
