+++
name = "organization-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_update"]

[match]
intent = "Updates the user's workspace settings. Different settings require different permission levels; most require the workspaceSettings admin permission."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "agentAutomationEnabled"
type = "Boolean"
description = "[INTERNAL] Whether the workspace has enabled agent automation."
required = false

[[inputs]]
name = "aiAddonEnabled"
type = "Boolean"
description = "[INTERNAL] Whether the workspace has enabled the AI add-on."
required = false

[[inputs]]
name = "aiDiscussionSummariesEnabled"
type = "Boolean"
description = "Whether the workspace has enabled AI discussion summaries for issues."
required = false

[[inputs]]
name = "aiTelemetryEnabled"
type = "Boolean"
description = "[INTERNAL] Whether the workspace has opted in to AI telemetry."
required = false

[[inputs]]
name = "aiThreadSummariesEnabled"
type = "Boolean"
description = "Whether the workspace has enabled resolved thread AI summaries."
required = false

[[inputs]]
name = "allowedAuthServices"
type = "String"
description = "List of services that are allowed to be used for login."
required = false

[[inputs]]
name = "allowedFileUploadContentTypes"
type = "String"
description = "Allowed file upload content types."
required = false

[[inputs]]
name = "authSettings"
type = "OrganizationAuthSettingsInput"
description = "The authentication settings for the workspace."
required = false

[[inputs]]
name = "codeIntelligenceEnabled"
type = "Boolean"
description = "[INTERNAL] Whether code intelligence is enabled for the workspace."
required = false

[[inputs]]
name = "codeIntelligenceRepository"
type = "String"
description = "[INTERNAL] GitHub repository in owner/repo format for code intelligence."
required = false

[[inputs]]
name = "codingAgentEnabled"
type = "Boolean"
description = "[INTERNAL] Whether the workspace has enabled Coding Sessions."
required = false

[[inputs]]
name = "codingAgentSettings"
type = "OrganizationCodingAgentSettingsInput"
description = "[Internal] Settings for Coding Sessions features."
required = false

[[inputs]]
name = "customersConfiguration"
type = "CustomersConfigurationInput"
description = "[INTERNAL] Configuration settings for the Customers feature."
required = false

[[inputs]]
name = "customersEnabled"
type = "Boolean"
description = "[INTERNAL] Whether the workspace is using customers."
required = false

[[inputs]]
name = "defaultFeedSummarySchedule"
type = "FeedSummarySchedule"
description = "Default schedule for how often feed summaries are generated."
required = false

[[inputs]]
name = "feedEnabled"
type = "Boolean"
description = "Whether the workspace has enabled the feed feature."
required = false

[[inputs]]
name = "fiscalYearStartMonth"
type = "Float"
description = "The month at which the fiscal year starts."
required = false

[[inputs]]
name = "generatedUpdatesEnabled"
type = "Boolean"
description = "[INTERNAL] Whether the workspace has enabled generated updates."
required = false

[[inputs]]
name = "gitBranchFormat"
type = "String"
description = "How git branches are formatted. If null, default formatting will be used."
required = false

[[inputs]]
name = "gitLinkbackDescriptionsEnabled"
type = "Boolean"
description = "Whether issue descriptions should be included in Git integration linkback messages."
required = false

[[inputs]]
name = "gitLinkbackMessagesEnabled"
type = "Boolean"
description = "Whether the Git integration linkback messages should be sent for private repositories."
required = false

[[inputs]]
name = "gitPublicLinkbackMessagesEnabled"
type = "Boolean"
description = "Whether the Git integration linkback messages should be sent for public repositories."
required = false

[[inputs]]
name = "hideNonPrimaryOrganizations"
type = "Boolean"
description = "Whether to hide other workspaces for new users signing up with email domains claimed by this organization."
required = false

[[inputs]]
name = "hipaaComplianceEnabled"
type = "Boolean"
description = "Whether HIPAA compliance is enabled for the workspace."
required = false

[[inputs]]
name = "initiativeUpdateReminderFrequencyInWeeks"
type = "Float"
description = "[ALPHA] The n-weekly frequency at which to prompt for initiative updates."
required = false

[[inputs]]
name = "initiativeUpdateRemindersDay"
type = "Day"
description = "[ALPHA] The day at which initiative updates are sent."
required = false

[[inputs]]
name = "initiativeUpdateRemindersHour"
type = "Float"
description = "[ALPHA] The hour at which initiative updates are sent."
required = false

[[inputs]]
name = "ipRestrictions"
type = "OrganizationIpRestrictionInput"
description = "IP restriction configurations controlling allowed access the workspace."
required = false

[[inputs]]
name = "linearAgentEnabled"
type = "Boolean"
description = "[Internal] Whether the workspace has enabled Linear Agent."
required = false

[[inputs]]
name = "linearAgentSettings"
type = "OrganizationLinearAgentSettingsInput"
description = "[Internal] Settings for Linear Agent features."
required = false

[[inputs]]
name = "logoUrl"
type = "String"
description = "The logo URL of the workspace."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the workspace."
required = false

[[inputs]]
name = "oauthAppReview"
type = "Boolean"
description = "Whether the workspace has opted for having to approve all OAuth applications for install."
required = false

[[inputs]]
name = "projectUpdateReminderFrequencyInWeeks"
type = "Float"
description = "The n-weekly frequency at which to prompt for project updates."
required = false

[[inputs]]
name = "projectUpdateRemindersDay"
type = "Day"
description = "The day at which project updates are sent."
required = false

[[inputs]]
name = "projectUpdateRemindersHour"
type = "Float"
description = "The hour at which project updates are sent."
required = false

[[inputs]]
name = "pullRequestTourEnabled"
type = "Boolean"
description = "Whether the workspace generates AI Pull Request guides for new pull requests."
required = false

[[inputs]]
name = "reducedPersonalInformation"
type = "Boolean"
description = "Whether the workspace has opted for reduced customer support attachment information."
required = false

[[inputs]]
name = "restrictAgentInvocationToMembers"
type = "Boolean"
description = "Whether agent invocation is restricted to full workspace members."
required = false

[[inputs]]
name = "roadmapEnabled"
type = "Boolean"
description = "Whether the workspace is using roadmap."
required = false

[[inputs]]
name = "securitySettings"
type = "OrganizationSecuritySettingsInput"
description = "The security settings for the workspace."
required = false

[[inputs]]
name = "slaEnabled"
type = "Boolean"
description = "Internal. Whether SLAs have been enabled for the workspace."
required = false

[[inputs]]
name = "slackAutoCreateProjectChannel"
type = "Boolean"
description = "[Internal] Whether to automatically create a Slack channel when a new project is created."
required = false

[[inputs]]
name = "slackProjectChannelIntegrationId"
type = "String"
description = "The ID of the Slack integration to use for auto-creating project channels."
required = false

[[inputs]]
name = "slackProjectChannelPrefix"
type = "String"
description = "The prefix to use for auto-created Slack project channels (p-, proj-, or project-)."
required = false

[[inputs]]
name = "slackProjectChannelsEnabled"
type = "Boolean"
description = "[Internal] Whether the Slack project channels feature is enabled for the workspace."
required = false

[[inputs]]
name = "themeSettings"
type = "OrganizationThemeSettingsInput"
description = "[ALPHA] Theme settings for the workspace."
required = false

[[inputs]]
name = "urlKey"
type = "String"
description = "The URL key of the workspace."
required = false

[[inputs]]
name = "workingDays"
type = "Float"
description = "[Internal] The list of working days. Sunday is 0, Monday is 1, etc."
required = false
+++
