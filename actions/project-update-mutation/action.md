+++
name = "project-update-mutation"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-update-mutation@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_update_mutation"]

[match]
intent = "Updates a project."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_update_mutation"
idempotent = false

[approval]
required = true

[[inputs]]
name = "canceledAt"
type = "DateTime"
description = "The time at which the project was canceled."
required = false

[[inputs]]
name = "color"
type = "String"
description = "The color of the project."
required = false

[[inputs]]
name = "completedAt"
type = "DateTime"
description = "The time at which the project was completed."
required = false

[[inputs]]
name = "content"
type = "String"
description = "The project content as markdown."
required = false

[[inputs]]
name = "convertedFromIssueId"
type = "String"
description = "The ID of the issue from which that project is created."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description for the project."
required = false

[[inputs]]
name = "frequencyResolution"
type = "FrequencyResolutionType"
description = "The resolution type for the update reminder frequency (e.g., weekly, biweekly)."
required = false

[[inputs]]
name = "icon"
type = "String"
description = "The icon of the project."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project to update. Also the identifier from the URL is accepted."
required = true

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
required = false

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
name = "projectUpdateRemindersPausedUntilAt"
type = "DateTime"
description = "The time until which project update reminders are paused. Set to null to resume reminders."
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
name = "sortOrder"
type = "Float"
description = "The sort order for the project in shared views."
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
required = false

[[inputs]]
name = "trashed"
type = "Boolean"
description = "Whether the project has been trashed. Set to true to trash, or null to restore."
required = false

[[inputs]]
name = "updateReminderFrequency"
type = "Float"
description = "The frequency at which to prompt for project updates. When not set, reminders are inherited from workspace settings."
required = false

[[inputs]]
name = "updateReminderFrequencyInWeeks"
type = "Float"
description = "The n-weekly frequency at which to prompt for project updates. When not set, reminders are inherited from workspace settings."
required = false

[[inputs]]
name = "updateRemindersDay"
type = "Day"
description = "The day of the week on which to prompt for project updates."
required = false

[[inputs]]
name = "updateRemindersHour"
type = "Int"
description = "The hour of the day (0-23) at which to prompt for project updates."
required = false
+++
