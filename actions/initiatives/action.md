+++
name = "initiatives"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiatives@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiatives"]

[match]
intent = "Returns all initiatives in the workspace, with optional filtering and sorting."

[[execute]]
id = "initiatives"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiatives"
idempotent = true

[[inputs]]
name = "activityType"
type = "StringComparator"
description = "Comparator for the initiative activity type."
required = false

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "ancestors"
type = "InitiativeCollectionFilter"
description = "Filters that the initiative must be an ancestor of."
required = false

[[inputs]]
name = "and"
type = "InitiativeFilter"
description = "Compound filters, all of which need to be matched by the initiative."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "completedAt"
type = "NullableDateComparator"
description = "Comparator for the initiative completed at date."
required = false

[[inputs]]
name = "createdAt"
type = "InitiativeCreatedAtSort"
description = "Sort by initiative creation date."
required = false

[[inputs]]
name = "createdAt"
type = "DateComparator"
description = "Comparator for the created at date."
required = false

[[inputs]]
name = "creator"
type = "NullableUserFilter"
description = "Filters that the initiative creator must satisfy."
required = false

[[inputs]]
name = "first"
type = "Int"
description = "The number of items to forward paginate (used with after). Defaults to 50."
required = false

[[inputs]]
name = "health"
type = "StringComparator"
description = "Comparator for the initiative health: onTrack, atRisk, offTrack"
required = false

[[inputs]]
name = "health"
type = "InitiativeHealthSort"
description = "Sort by initiative health status."
required = false

[[inputs]]
name = "healthUpdatedAt"
type = "InitiativeHealthUpdatedAtSort"
description = "Sort by initiative health update date."
required = false

[[inputs]]
name = "healthWithAge"
type = "StringComparator"
description = "Comparator for the initiative health (with age): onTrack, atRisk, offTrack, outdated, noUpdate"
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
name = "initiativeUpdates"
type = "InitiativeUpdatesCollectionFilter"
description = "Filters that the initiative updates must satisfy."
required = false

[[inputs]]
name = "labels"
type = "InitiativeLabelCollectionFilter"
description = "[Internal] Filters that the initiative labels must satisfy."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "manual"
type = "InitiativeManualSort"
description = "Sort by manual order."
required = false

[[inputs]]
name = "name"
type = "InitiativeNameSort"
description = "Sort by initiative name."
required = false

[[inputs]]
name = "name"
type = "StringComparator"
description = "Comparator for the initiative name."
required = false

[[inputs]]
name = "or"
type = "InitiativeFilter"
description = "Compound filters, one of which need to be matched by the initiative."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "owner"
type = "NullableUserFilter"
description = "Filters that the initiative owner must satisfy."
required = false

[[inputs]]
name = "owner"
type = "InitiativeOwnerSort"
description = "Sort by initiative owner name."
required = false

[[inputs]]
name = "slugId"
type = "StringComparator"
description = "Comparator for the initiative slug ID."
required = false

[[inputs]]
name = "startedAt"
type = "NullableDateComparator"
description = "Comparator for the initiative started at date."
required = false

[[inputs]]
name = "status"
type = "StringComparator"
description = "Comparator for the initiative status: Planned, Active, Completed"
required = false

[[inputs]]
name = "targetDate"
type = "NullableDateComparator"
description = "Comparator for the initiative target date."
required = false

[[inputs]]
name = "targetDate"
type = "InitiativeTargetDateSort"
description = "Sort by initiative target date."
required = false

[[inputs]]
name = "teams"
type = "TeamCollectionFilter"
description = "Filters that the initiative teams must satisfy."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false

[[inputs]]
name = "updatedAt"
type = "InitiativeUpdatedAtSort"
description = "Sort by initiative update date."
required = false
+++
