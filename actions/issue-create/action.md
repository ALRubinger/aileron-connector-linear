+++
name = "issue-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["create_issue"]

[match]
intent = "Creates a new issue."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "assigneeId"
type = "String"
description = "The identifier of the user to assign the issue to."
required = false

[[inputs]]
name = "completedAt"
type = "DateTime"
description = "The time at which the issue was completed (e.g. if importing from another system). Must be a time in the past and after createdAt. Cannot be provided with an incompatible workflow state."
required = false

[[inputs]]
name = "createAsUser"
type = "String"
description = "Create issue as a user with the provided name. This option is only available to OAuth applications creating issues in `actor=app` mode."
required = false

[[inputs]]
name = "createdAt"
type = "DateTime"
description = "The time at which the issue was created (e.g. if importing from another system). Must be a time in the past. If none is provided, the backend will generate the time as now."
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
name = "displayIconUrl"
type = "String"
description = "Provide an external user avatar URL. Can only be used in conjunction with the `createAsUser` options. This option is only available to OAuth applications creating comments in `actor=app` mode."
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
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "inheritsSharedAccess"
type = "Boolean"
description = "[Internal] Whether this issue should inherit shared access from its parent issue. Set to false to opt out of automatic shared access inheritance when creating a sub-issue."
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
name = "preserveSortOrderOnCreate"
type = "Boolean"
description = "Whether the passed sort order should be preserved."
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
name = "referenceCommentId"
type = "String"
description = "The comment the issue is referencing."
required = false

[[inputs]]
name = "releaseIds"
type = "String"
description = "The identifiers of the releases to associate with this issue."
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
name = "sortOrder"
type = "Float"
description = "The position of the issue related to other issues."
required = false

[[inputs]]
name = "sourceCommentId"
type = "String"
description = "The comment the issue is created from."
required = false

[[inputs]]
name = "sourcePullRequestCommentId"
type = "String"
description = "[Internal] The pull request comment the issue is created from."
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
required = true

[[inputs]]
name = "templateId"
type = "String"
description = "The identifier of a template the issue should be created from. If other values are provided in the input, they will override template values."
required = false

[[inputs]]
name = "title"
type = "String"
description = "The title of the issue."
required = false

[[inputs]]
name = "useDefaultTemplate"
type = "Boolean"
description = "Whether to use the default template for the team. When set to true, the default template of this team based on user's membership will be applied."
required = false
+++
