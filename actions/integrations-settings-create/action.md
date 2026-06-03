+++
name = "integrations-settings-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integrations-settings-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integrations_settings_create"]

[match]
intent = "Creates new Slack notification settings for a team, project, initiative, or custom view."

[[execute]]
id = "integrations"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integrations_settings_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "contextViewType"
type = "ContextViewType"
description = "The type of view to which the integration settings context is associated with."
required = false

[[inputs]]
name = "customViewId"
type = "String"
description = "The identifier of the custom view to create settings for."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "The identifier of the initiative to create settings for."
required = false

[[inputs]]
name = "microsoftTeamsProjectUpdateCreated"
type = "Boolean"
description = "Whether to send a Microsoft Teams message when a project update is created."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "The identifier of the project to create settings for."
required = false

[[inputs]]
name = "slackInitiativeUpdateCreated"
type = "Boolean"
description = "Whether to send a Slack message when an initiative update is created."
required = false

[[inputs]]
name = "slackIssueAddedToTriage"
type = "Boolean"
description = "Whether to send a Slack message when a new issue is added to triage."
required = false

[[inputs]]
name = "slackIssueAddedToView"
type = "Boolean"
description = "Whether to send a Slack message when an issue is added to a view."
required = false

[[inputs]]
name = "slackIssueCreated"
type = "Boolean"
description = "Whether to send a Slack message when a new issue is created for the project or the team."
required = false

[[inputs]]
name = "slackIssueNewComment"
type = "Boolean"
description = "Whether to send a Slack message when a comment is created on any of the project or team's issues."
required = false

[[inputs]]
name = "slackIssueSlaBreached"
type = "Boolean"
description = "Whether to receive notification when an SLA has breached on Slack."
required = false

[[inputs]]
name = "slackIssueSlaHighRisk"
type = "Boolean"
description = "Whether to send a Slack message when an SLA is at high risk."
required = false

[[inputs]]
name = "slackIssueStatusChangedAll"
type = "Boolean"
description = "Whether to send a Slack message when any of the project or team's issues has a change in status."
required = false

[[inputs]]
name = "slackIssueStatusChangedDone"
type = "Boolean"
description = "Whether to send a Slack message when any of the project or team's issues change to completed or canceled."
required = false

[[inputs]]
name = "slackProjectUpdateCreated"
type = "Boolean"
description = "Whether to send a Slack message when a project update is created."
required = false

[[inputs]]
name = "slackProjectUpdateCreatedToTeam"
type = "Boolean"
description = "Whether to send a Slack message when a project update is created to team channels."
required = false

[[inputs]]
name = "slackProjectUpdateCreatedToWorkspace"
type = "Boolean"
description = "Whether to send a Slack message when a project update is created to workspace channel."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier of the team to create settings for."
required = false
+++
