+++
name = "issue-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_update"]

[match]
intent = "Updates an issue."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "addedLabelIds"
type = "String"
description = "The identifiers of the issue labels to be added to this issue."
required = false

[[inputs]]
name = "addedReleaseIds"
type = "String"
description = "The identifiers of the releases to be added to this issue."
required = false

[[inputs]]
name = "assigneeId"
type = "String"
description = "The identifier of the user to assign the issue to."
required = false

[[inputs]]
name = "autoClosedByParentClosing"
type = "Boolean"
description = "Whether the issue was automatically closed because its parent issue was closed."
required = false

[[inputs]]
name = "cycleId"
type = "String"
description = "The cycle associated with the issue."
required = false

[[inputs]]
name = "delegateId"
type = "String"
description = "The identifier of the agent user to delegate the issue to."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The issue description in markdown format."
required = false

[[inputs]]
name = "descriptionData"
type = "JSON"
description = "[Internal] The issue description as a Prosemirror document."
required = false

[[inputs]]
name = "dueDate"
type = "TimelessDate"
description = "The date at which the issue is due."
required = false

[[inputs]]
name = "estimate"
type = "Int"
description = "The estimated complexity of the issue."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to update."
required = true

[[inputs]]
name = "inheritsSharedAccess"
type = "Boolean"
description = "Whether this issue should inherit shared access from its parent issue."
required = false

[[inputs]]
name = "labelIds"
type = "String"
description = "The identifiers of the issue labels associated with this ticket."
required = false

[[inputs]]
name = "lastAppliedTemplateId"
type = "String"
description = "The ID of the last template applied to the issue."
required = false

[[inputs]]
name = "parentId"
type = "String"
description = "The identifier of the parent issue. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = false

[[inputs]]
name = "priority"
type = "Int"
description = "The priority of the issue. 0 = No priority, 1 = Urgent, 2 = High, 3 = Medium, 4 = Low."
required = false

[[inputs]]
name = "prioritySortOrder"
type = "Float"
description = "The position of the issue related to other issues, when ordered by priority."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "The project associated with the issue."
required = false

[[inputs]]
name = "projectMilestoneId"
type = "String"
description = "The project milestone associated with the issue."
required = false

[[inputs]]
name = "releaseIds"
type = "String"
description = "The identifiers of the releases associated with this issue."
required = false

[[inputs]]
name = "removedLabelIds"
type = "String"
description = "The identifiers of the issue labels to be removed from this issue."
required = false

[[inputs]]
name = "removedReleaseIds"
type = "String"
description = "The identifiers of the releases to be removed from this issue."
required = false

[[inputs]]
name = "slaBreachesAt"
type = "DateTime"
description = "[Internal] The time at which an issue will be considered in breach of SLA."
required = false

[[inputs]]
name = "slaStartedAt"
type = "DateTime"
description = "[Internal] The time at which the issue's SLA was started."
required = false

[[inputs]]
name = "slaType"
type = "SLADayCountType"
description = "The SLA day count type for the issue. Whether SLA should be business days only or calendar days (default)."
required = false

[[inputs]]
name = "snoozedById"
type = "String"
description = "The identifier of the user who snoozed the issue."
required = false

[[inputs]]
name = "snoozedUntilAt"
type = "DateTime"
description = "The time until which the issue will be snoozed in Triage view."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The position of the issue related to other issues."
required = false

[[inputs]]
name = "stateId"
type = "String"
description = "The team state of the issue."
required = false

[[inputs]]
name = "subIssueSortOrder"
type = "Float"
description = "The position of the issue in parent's sub-issue list."
required = false

[[inputs]]
name = "subscriberIds"
type = "String"
description = "The identifiers of the users subscribing to this ticket."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier of the team associated with the issue."
required = false

[[inputs]]
name = "title"
type = "String"
description = "The issue title."
required = false

[[inputs]]
name = "trashed"
type = "Boolean"
description = "Whether the issue has been trashed."
required = false
+++
