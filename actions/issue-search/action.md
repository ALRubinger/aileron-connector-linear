+++
name = "issue-search"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-search@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_search"]

[match]
intent = "[DEPRECATED] Search issues. This endpoint is deprecated and will be removed in the future – use `searchIssues` instead."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_search"
idempotent = true

[[inputs]]
name = "accumulatedStateUpdatedAt"
type = "NullableDateComparator"
description = "[Internal] Comparator for the issue's accumulatedStateUpdatedAt date."
required = false

[[inputs]]
name = "activity"
type = "ActivityCollectionFilter"
description = "Filters that the issue's activities must satisfy."
required = false

[[inputs]]
name = "addedToCycleAt"
type = "NullableDateComparator"
description = "Comparator for the issues added to cycle at date."
required = false

[[inputs]]
name = "addedToCyclePeriod"
type = "CyclePeriodComparator"
description = "Comparator for the period when issue was added to a cycle."
required = false

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "ageTime"
type = "NullableDurationComparator"
description = "[Internal] Age (created -> now) comparator, defined if the issue is still open."
required = false

[[inputs]]
name = "and"
type = "IssueFilter"
description = "Compound filters, all of which need to be matched by the issue."
required = false

[[inputs]]
name = "archivedAt"
type = "NullableDateComparator"
description = "Comparator for the issues archived at date."
required = false

[[inputs]]
name = "assignee"
type = "NullableUserFilter"
description = "Filters that the issues assignee must satisfy."
required = false

[[inputs]]
name = "attachments"
type = "AttachmentCollectionFilter"
description = "Filters that the issues attachments must satisfy."
required = false

[[inputs]]
name = "autoArchivedAt"
type = "NullableDateComparator"
description = "Comparator for the issues auto archived at date."
required = false

[[inputs]]
name = "autoClosedAt"
type = "NullableDateComparator"
description = "Comparator for the issues auto closed at date."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "canceledAt"
type = "NullableDateComparator"
description = "Comparator for the issues canceled at date."
required = false

[[inputs]]
name = "children"
type = "IssueCollectionFilter"
description = "Filters that the child issues must satisfy."
required = false

[[inputs]]
name = "comments"
type = "CommentCollectionFilter"
description = "Filters that the issues comments must satisfy."
required = false

[[inputs]]
name = "completedAt"
type = "NullableDateComparator"
description = "Comparator for the issues completed at date."
required = false

[[inputs]]
name = "createdAt"
type = "DateComparator"
description = "Comparator for the created at date."
required = false

[[inputs]]
name = "creator"
type = "NullableUserFilter"
description = "Filters that the issues creator must satisfy."
required = false

[[inputs]]
name = "customerCount"
type = "NumberComparator"
description = "Count of customers"
required = false

[[inputs]]
name = "customerImportantCount"
type = "NumberComparator"
description = "Count of important customers"
required = false

[[inputs]]
name = "cycle"
type = "NullableCycleFilter"
description = "Filters that the issues cycle must satisfy."
required = false

[[inputs]]
name = "cycleTime"
type = "NullableDurationComparator"
description = "[Internal] Cycle time (started -> completed) comparator."
required = false

[[inputs]]
name = "delegate"
type = "NullableUserFilter"
description = "Filters that the issue's delegated agent must satisfy."
required = false

[[inputs]]
name = "description"
type = "NullableStringComparator"
description = "Comparator for the issues description."
required = false

[[inputs]]
name = "dueDate"
type = "NullableTimelessDateComparator"
description = "Comparator for the issues due date."
required = false

[[inputs]]
name = "estimate"
type = "EstimateComparator"
description = "Comparator for the issues estimate."
required = false

[[inputs]]
name = "first"
type = "Int"
description = "The number of items to forward paginate (used with after). Defaults to 50."
required = false

[[inputs]]
name = "hasActiveAgentSessions"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues bucketed as having active agent sessions, no closed agent pull requests, and no merged agent pull requests."
required = false

[[inputs]]
name = "hasBlockedByRelations"
type = "RelationExistsComparator"
description = "Comparator for filtering issues which are blocked."
required = false

[[inputs]]
name = "hasBlockingRelations"
type = "RelationExistsComparator"
description = "Comparator for filtering issues which are blocking."
required = false

[[inputs]]
name = "hasDismissedAgentSessions"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues bucketed as having all agent sessions dismissed or closed, and no merged agent pull requests."
required = false

[[inputs]]
name = "hasDuplicateRelations"
type = "RelationExistsComparator"
description = "Comparator for filtering issues which are duplicates."
required = false

[[inputs]]
name = "hasErroredAgentSessions"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues bucketed as having an errored agent session, no active sessions, and no merged agent pull requests."
required = false

[[inputs]]
name = "hasMergedAgentPullRequests"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues which have agent-session-linked pull requests that were merged."
required = false

[[inputs]]
name = "hasRelatedRelations"
type = "RelationExistsComparator"
description = "Comparator for filtering issues with relations."
required = false

[[inputs]]
name = "hasSharedUsers"
type = "RelationExistsComparator"
description = "Comparator for filtering issues which have been shared with users outside of the team."
required = false

[[inputs]]
name = "hasSuggestedAssignees"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues which have suggested assignees."
required = false

[[inputs]]
name = "hasSuggestedLabels"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues which have suggested labels."
required = false

[[inputs]]
name = "hasSuggestedProjects"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues which have suggested projects."
required = false

[[inputs]]
name = "hasSuggestedRelatedIssues"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues which have suggested related issues."
required = false

[[inputs]]
name = "hasSuggestedSimilarIssues"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues which have suggested similar issues."
required = false

[[inputs]]
name = "hasSuggestedTeams"
type = "RelationExistsComparator"
description = "[Internal] Comparator for filtering issues which have suggested teams."
required = false

[[inputs]]
name = "id"
type = "IssueIDComparator"
description = "Comparator for the identifier."
required = false

[[inputs]]
name = "includeArchived"
type = "Boolean"
description = "Should archived resources be included (default: false)"
required = false

[[inputs]]
name = "labels"
type = "IssueLabelCollectionFilter"
description = "Filters that issue labels must satisfy."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "lastAppliedTemplate"
type = "NullableTemplateFilter"
description = "Filters that the last applied template must satisfy."
required = false

[[inputs]]
name = "leadTime"
type = "NullableDurationComparator"
description = "[Internal] Lead time (created -> completed) comparator."
required = false

[[inputs]]
name = "needs"
type = "CustomerNeedCollectionFilter"
description = "Filters that the issue's customer needs must satisfy."
required = false

[[inputs]]
name = "number"
type = "NumberComparator"
description = "Comparator for the issues number."
required = false

[[inputs]]
name = "or"
type = "IssueFilter"
description = "Compound filters, one of which need to be matched by the issue."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "parent"
type = "NullableIssueFilter"
description = "Filters that the issue parent must satisfy."
required = false

[[inputs]]
name = "priority"
type = "NullableNumberComparator"
description = "Comparator for the issues priority. 0 = No priority, 1 = Urgent, 2 = High, 3 = Medium, 4 = Low."
required = false

[[inputs]]
name = "project"
type = "NullableProjectFilter"
description = "Filters that the issues project must satisfy."
required = false

[[inputs]]
name = "projectMilestone"
type = "NullableProjectMilestoneFilter"
description = "Filters that the issues project milestone must satisfy."
required = false

[[inputs]]
name = "query"
type = "String"
description = "[Deprecated] Search string to look for."
required = false

[[inputs]]
name = "reactions"
type = "ReactionCollectionFilter"
description = "Filters that the issues reactions must satisfy."
required = false

[[inputs]]
name = "recurringIssueTemplate"
type = "NullableTemplateFilter"
description = "[ALPHA] Filters that the recurring issue template must satisfy."
required = false

[[inputs]]
name = "releases"
type = "ReleaseCollectionFilter"
description = "Filters that the issue's releases must satisfy."
required = false

[[inputs]]
name = "searchableContent"
type = "ContentComparator"
description = "[Internal] Comparator for the issues content."
required = false

[[inputs]]
name = "sharedWith"
type = "UserCollectionFilter"
description = "Filters that users the issue has been shared with must satisfy."
required = false

[[inputs]]
name = "slaBreachesAt"
type = "NullableDateComparator"
description = "Comparator for the issue's SLA breach date."
required = false

[[inputs]]
name = "slaStatus"
type = "SlaStatusComparator"
description = "Comparator for the issues sla status."
required = false

[[inputs]]
name = "snoozedBy"
type = "NullableUserFilter"
description = "Filters that the issues snoozer must satisfy."
required = false

[[inputs]]
name = "snoozedUntilAt"
type = "NullableDateComparator"
description = "Comparator for the issues snoozed until date."
required = false

[[inputs]]
name = "sourceMetadata"
type = "SourceMetadataComparator"
description = "Filters that the source must satisfy."
required = false

[[inputs]]
name = "startedAt"
type = "NullableDateComparator"
description = "Comparator for the issues started at date."
required = false

[[inputs]]
name = "state"
type = "WorkflowStateFilter"
description = "Filters that the issues state must satisfy."
required = false

[[inputs]]
name = "subscribers"
type = "UserCollectionFilter"
description = "Filters that issue subscribers must satisfy."
required = false

[[inputs]]
name = "suggestions"
type = "IssueSuggestionCollectionFilter"
description = "[Internal] Filters that the issue's suggestions must satisfy."
required = false

[[inputs]]
name = "team"
type = "TeamFilter"
description = "Filters that the issues team must satisfy."
required = false

[[inputs]]
name = "title"
type = "StringComparator"
description = "Comparator for the issues title."
required = false

[[inputs]]
name = "triageTime"
type = "NullableDurationComparator"
description = "[Internal] Triage time (entered triaged -> triaged) comparator."
required = false

[[inputs]]
name = "triagedAt"
type = "NullableDateComparator"
description = "Comparator for the issues triaged at date."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
