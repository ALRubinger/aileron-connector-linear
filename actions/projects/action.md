+++
name = "projects"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/projects@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["projects"]

[match]
intent = "Returns all projects in the workspace, with optional filtering and sorting."

[[execute]]
id = "projects"
connector = "github://ALRubinger/aileron-connector-linear"
op = "projects"
idempotent = true

[[inputs]]
name = "accessibleTeams"
type = "TeamCollectionFilter"
description = "Filters that the project's team must satisfy."
required = false

[[inputs]]
name = "activityType"
type = "StringComparator"
description = "[ALPHA] Comparator for the project activity type: buzzin, active, some, none"
required = false

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "ProjectFilter"
description = "Compound filters, all of which need to be matched by the project."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "canceledAt"
type = "NullableDateComparator"
description = "Comparator for the project cancelation date."
required = false

[[inputs]]
name = "completedAt"
type = "NullableDateComparator"
description = "Comparator for the project completion date."
required = false

[[inputs]]
name = "completedProjectMilestones"
type = "ProjectMilestoneCollectionFilter"
description = "Filters that the project's completed milestones must satisfy."
required = false

[[inputs]]
name = "createdAt"
type = "ProjectCreatedAtSort"
description = "Sort by project creation date"
required = false

[[inputs]]
name = "createdAt"
type = "DateComparator"
description = "Comparator for the created at date."
required = false

[[inputs]]
name = "creator"
type = "UserFilter"
description = "Filters that the projects creator must satisfy."
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
name = "first"
type = "Int"
description = "The number of items to forward paginate (used with after). Defaults to 50."
required = false

[[inputs]]
name = "hasBlockedByRelations"
type = "RelationExistsComparator"
description = "Comparator for filtering projects which are blocked."
required = false

[[inputs]]
name = "hasBlockingRelations"
type = "RelationExistsComparator"
description = "Comparator for filtering projects which are blocking."
required = false

[[inputs]]
name = "hasDependedOnByRelations"
type = "RelationExistsComparator"
description = "[Deprecated] Comparator for filtering projects which this is depended on by."
required = false

[[inputs]]
name = "hasDependsOnRelations"
type = "RelationExistsComparator"
description = "[Deprecated]Comparator for filtering projects which this depends on."
required = false

[[inputs]]
name = "hasRelatedRelations"
type = "RelationExistsComparator"
description = "Comparator for filtering projects with relations."
required = false

[[inputs]]
name = "hasViolatedRelations"
type = "RelationExistsComparator"
description = "Comparator for filtering projects with violated dependencies."
required = false

[[inputs]]
name = "health"
type = "StringComparator"
description = "Comparator for the project health: onTrack, atRisk, offTrack"
required = false

[[inputs]]
name = "health"
type = "ProjectHealthSort"
description = "Sort by project health status."
required = false

[[inputs]]
name = "healthWithAge"
type = "StringComparator"
description = "Comparator for the project health (with age): onTrack, atRisk, offTrack, outdated, noUpdate"
required = false

[[inputs]]
name = "id"
type = "IDComparator"
description = "Comparator for the identifier."
required = false

[[inputs]]
name = "includeArchived"
type = "Boolean"
description = "Should archived resources be included (default: false)"
required = false

[[inputs]]
name = "initiatives"
type = "InitiativeCollectionFilter"
description = "Filters that the projects initiatives must satisfy."
required = false

[[inputs]]
name = "issues"
type = "IssueCollectionFilter"
description = "Filters that the projects issues must satisfy."
required = false

[[inputs]]
name = "labels"
type = "ProjectLabelCollectionFilter"
description = "Filters that project labels must satisfy."
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
name = "lead"
type = "NullableUserFilter"
description = "Filters that the projects lead must satisfy."
required = false

[[inputs]]
name = "lead"
type = "ProjectLeadSort"
description = "Sort by project lead name."
required = false

[[inputs]]
name = "manual"
type = "ProjectManualSort"
description = "Sort by manual order"
required = false

[[inputs]]
name = "members"
type = "UserCollectionFilter"
description = "Filters that the projects members must satisfy."
required = false

[[inputs]]
name = "name"
type = "ProjectNameSort"
description = "Sort by project name"
required = false

[[inputs]]
name = "name"
type = "StringComparator"
description = "Comparator for the project name."
required = false

[[inputs]]
name = "needs"
type = "CustomerNeedCollectionFilter"
description = "Filters that the project's customer needs must satisfy."
required = false

[[inputs]]
name = "nextProjectMilestone"
type = "ProjectMilestoneFilter"
description = "Filters that the project's next milestone must satisfy."
required = false

[[inputs]]
name = "or"
type = "ProjectFilter"
description = "Compound filters, one of which need to be matched by the project."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "priority"
type = "ProjectPrioritySort"
description = "Sort by project priority"
required = false

[[inputs]]
name = "priority"
type = "NullableNumberComparator"
description = "Comparator for the projects priority."
required = false

[[inputs]]
name = "projectMilestones"
type = "ProjectMilestoneCollectionFilter"
description = "Filters that the project's milestones must satisfy."
required = false

[[inputs]]
name = "projectUpdates"
type = "ProjectUpdatesCollectionFilter"
description = "Comparator for the project updates."
required = false

[[inputs]]
name = "roadmaps"
type = "RoadmapCollectionFilter"
description = "Filters that the projects roadmaps must satisfy."
required = false

[[inputs]]
name = "searchableContent"
type = "ContentComparator"
description = "[Internal] Comparator for the project's content."
required = false

[[inputs]]
name = "slugId"
type = "StringComparator"
description = "Comparator for the project slug ID."
required = false

[[inputs]]
name = "startDate"
type = "StartDateSort"
description = "Sort by project start date"
required = false

[[inputs]]
name = "startDate"
type = "NullableDateComparator"
description = "Comparator for the project start date."
required = false

[[inputs]]
name = "startedAt"
type = "NullableDateComparator"
description = "Comparator for the project started date (when it was moved to an \"In Progress\" status)."
required = false

[[inputs]]
name = "state"
type = "StringComparator"
description = "[DEPRECATED] Comparator for the project state."
required = false

[[inputs]]
name = "status"
type = "ProjectStatusFilter"
description = "Filters that the project's status must satisfy."
required = false

[[inputs]]
name = "status"
type = "ProjectStatusSort"
description = "Sort by project status"
required = false

[[inputs]]
name = "targetDate"
type = "NullableDateComparator"
description = "Comparator for the project target date."
required = false

[[inputs]]
name = "targetDate"
type = "TargetDateSort"
description = "Sort by project target date"
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false

[[inputs]]
name = "updatedAt"
type = "ProjectUpdatedAtSort"
description = "Sort by project update date"
required = false
+++
