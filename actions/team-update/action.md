+++
name = "team-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_update"]

[match]
intent = "Updates a team's settings, properties, or configuration. Requires team owner or workspace admin permissions for most changes."

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "aiDiscussionSummariesEnabled"
type = "Boolean"
description = "Whether to enable AI discussion summaries for issues."
required = false

[[inputs]]
name = "aiThreadSummariesEnabled"
type = "Boolean"
description = "Whether to enable resolved thread AI summaries."
required = false

[[inputs]]
name = "allMembersCanJoin"
type = "Boolean"
description = "Whether all members in the workspace can join the team. Only used for public teams."
required = false

[[inputs]]
name = "autoArchivePeriod"
type = "Float"
description = "Period after which closed (completed, canceled, or duplicate) issues are automatically archived, in months."
required = false

[[inputs]]
name = "autoCloseChildIssues"
type = "Boolean"
description = "Whether to automatically close all sub-issues when a parent issue in this team is closed."
required = false

[[inputs]]
name = "autoCloseParentIssues"
type = "Boolean"
description = "Whether to automatically close a parent issue in this team if all its sub-issues are closed."
required = false

[[inputs]]
name = "autoClosePeriod"
type = "Float"
description = "Period after which issues are automatically closed, in months."
required = false

[[inputs]]
name = "autoCloseStateId"
type = "String"
description = "The canceled workflow state which auto closed issues will be set to."
required = false

[[inputs]]
name = "color"
type = "String"
description = "The color of the team."
required = false

[[inputs]]
name = "cycleCooldownTime"
type = "Int"
description = "The cooldown time after each cycle in weeks."
required = false

[[inputs]]
name = "cycleDuration"
type = "Int"
description = "The duration of each cycle in weeks."
required = false

[[inputs]]
name = "cycleEnabledStartDate"
type = "DateTime"
description = "The time at which to begin cycles."
required = false

[[inputs]]
name = "cycleIssueAutoAssignCompleted"
type = "Boolean"
description = "Auto assign completed issues to current active cycle setting."
required = false

[[inputs]]
name = "cycleIssueAutoAssignStarted"
type = "Boolean"
description = "Auto assign started issues to current active cycle setting."
required = false

[[inputs]]
name = "cycleLockToActive"
type = "Boolean"
description = "Only allow issues with cycles in Active Issues."
required = false

[[inputs]]
name = "cycleStartDay"
type = "Float"
description = "The day of the week that a new cycle starts."
required = false

[[inputs]]
name = "cyclesEnabled"
type = "Boolean"
description = "Whether the team uses cycles."
required = false

[[inputs]]
name = "defaultIssueEstimate"
type = "Float"
description = "What to use as an default estimate for unestimated issues."
required = false

[[inputs]]
name = "defaultIssueStateId"
type = "String"
description = "Default status for newly created issues."
required = false

[[inputs]]
name = "defaultProjectTemplateId"
type = "String"
description = "The identifier of the default project template of this team."
required = false

[[inputs]]
name = "defaultTemplateForMembersId"
type = "String"
description = "The identifier of the default template for members of this team."
required = false

[[inputs]]
name = "defaultTemplateForNonMembersId"
type = "String"
description = "The identifier of the default template for non-members of this team."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the team."
required = false

[[inputs]]
name = "groupIssueHistory"
type = "Boolean"
description = "Whether to group recent issue history entries."
required = false

[[inputs]]
name = "handleSubTeamsOnRetirement"
type = "TeamRetirementSubTeamHandling"
description = "[Internal] How to handle sub-teams when retiring. Required if the team has active sub-teams."
required = false

[[inputs]]
name = "icon"
type = "String"
description = "The icon of the team."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the team to update."
required = true

[[inputs]]
name = "inheritIssueEstimation"
type = "Boolean"
description = "Whether the team should inherit estimation settings from its parent. Only applies to sub-teams."
required = false

[[inputs]]
name = "inheritProductIntelligenceScope"
type = "Boolean"
description = "[Internal] Whether the team should inherit its product intelligence scope from its parent. Only applies to sub-teams."
required = false

[[inputs]]
name = "inheritSlackAutoCreateProjectChannel"
type = "Boolean"
description = "[Internal] Whether the team should inherit its Slack auto-create project channel setting from its parent. Only applies to sub-teams."
required = false

[[inputs]]
name = "inheritWorkflowStatuses"
type = "Boolean"
description = "[Internal] Whether the team should inherit workflow statuses from its parent."
required = false

[[inputs]]
name = "issueEstimationAllowZero"
type = "Boolean"
description = "Whether to allow zeros in issues estimates."
required = false

[[inputs]]
name = "issueEstimationExtended"
type = "Boolean"
description = "Whether to add additional points to the estimate scale."
required = false

[[inputs]]
name = "issueEstimationType"
type = "String"
description = "The issue estimation type to use. Must be one of \"notUsed\", \"exponential\", \"fibonacci\", \"linear\", \"tShirt\"."
required = false

[[inputs]]
name = "issueLabels"
type = "JSONObject"
description = "Mapping of the IssueLabel ID to the new IssueLabel name."
required = false

[[inputs]]
name = "issueSharingEnabled"
type = "Boolean"
description = "Whether issue sharing is enabled for this team."
required = false

[[inputs]]
name = "joinByDefault"
type = "Boolean"
description = "Whether new users should join this team by default. Mutation restricted to workspace admins or owners!"
required = false

[[inputs]]
name = "key"
type = "String"
description = "The key of the team."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the team."
required = false

[[inputs]]
name = "parentId"
type = "String"
description = "The parent team ID."
required = false

[[inputs]]
name = "private"
type = "Boolean"
description = "Whether the team is private or not."
required = false

[[inputs]]
name = "productIntelligenceScope"
type = "ProductIntelligenceScope"
description = "[Internal] The scope of product intelligence suggestion data for the team."
required = false

[[inputs]]
name = "requirePriorityToLeaveTriage"
type = "Boolean"
description = "Whether an issue needs to have a priority set before leaving triage."
required = false

[[inputs]]
name = "retiredAt"
type = "DateTime"
description = "When the team was retired."
required = false

[[inputs]]
name = "scimGroupName"
type = "String"
description = "The SCIM group name for the team."
required = false

[[inputs]]
name = "scimManaged"
type = "Boolean"
description = "Whether the team is managed by SCIM integration. Mutation restricted to workspace admins or owners and only unsetting is allowed!"
required = false

[[inputs]]
name = "securitySettings"
type = "TeamSecuritySettingsInput"
description = "The security settings for the team."
required = false

[[inputs]]
name = "setIssueSortOrderOnStateChange"
type = "String"
description = "Whether to move issues to bottom of the column when changing state."
required = false

[[inputs]]
name = "slackAutoCreateProjectChannel"
type = "Boolean"
description = "[Internal] Whether to automatically create a Slack channel when a new project is created in this team."
required = false

[[inputs]]
name = "slackIssueComments"
type = "Boolean"
description = "Whether to send new issue comment notifications to Slack."
required = false

[[inputs]]
name = "slackIssueStatuses"
type = "Boolean"
description = "Whether to send issue status update notifications to Slack."
required = false

[[inputs]]
name = "slackNewIssue"
type = "Boolean"
description = "Whether to send new issue notifications to Slack."
required = false

[[inputs]]
name = "timezone"
type = "String"
description = "The timezone of the team."
required = false

[[inputs]]
name = "triageEnabled"
type = "Boolean"
description = "Whether triage mode is enabled for the team."
required = false

[[inputs]]
name = "upcomingCycleCount"
type = "Float"
description = "How many upcoming cycles to create."
required = false

[[inputs]]
name = "workflowStates"
type = "JSONObject"
description = "Mapping of the WorkflowState ID to the new WorkflowState ID."
required = true
+++
