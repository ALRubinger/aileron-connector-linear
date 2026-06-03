+++
name = "cycles"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/cycles@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["cycles"]

[match]
intent = "All cycles accessible to the user."

[[execute]]
id = "cycles"
connector = "github://ALRubinger/aileron-connector-linear"
op = "cycles"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "CycleFilter"
description = "Compound filters, all of which need to be matched by the cycle."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "completedAt"
type = "DateComparator"
description = "Comparator for the cycle completed at date."
required = false

[[inputs]]
name = "createdAt"
type = "DateComparator"
description = "Comparator for the created at date."
required = false

[[inputs]]
name = "endsAt"
type = "DateComparator"
description = "Comparator for the cycle ends at date."
required = false

[[inputs]]
name = "first"
type = "Int"
description = "The number of items to forward paginate (used with after). Defaults to 50."
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
name = "inheritedFromId"
type = "IDComparator"
description = "Comparator for the inherited cycle ID."
required = false

[[inputs]]
name = "isActive"
type = "BooleanComparator"
description = "Comparator for the filtering active cycle."
required = false

[[inputs]]
name = "isFuture"
type = "BooleanComparator"
description = "Comparator for the filtering future cycles."
required = false

[[inputs]]
name = "isInCooldown"
type = "BooleanComparator"
description = "Comparator for filtering for whether the cycle is currently in cooldown."
required = false

[[inputs]]
name = "isNext"
type = "BooleanComparator"
description = "Comparator for the filtering next cycle."
required = false

[[inputs]]
name = "isPast"
type = "BooleanComparator"
description = "Comparator for the filtering past cycles."
required = false

[[inputs]]
name = "isPrevious"
type = "BooleanComparator"
description = "Comparator for the filtering previous cycle."
required = false

[[inputs]]
name = "issues"
type = "IssueCollectionFilter"
description = "Filters that the cycles issues must satisfy."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "name"
type = "StringComparator"
description = "Comparator for the cycle name."
required = false

[[inputs]]
name = "number"
type = "NumberComparator"
description = "Comparator for the cycle number."
required = false

[[inputs]]
name = "or"
type = "CycleFilter"
description = "Compound filters, one of which need to be matched by the cycle."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "startsAt"
type = "DateComparator"
description = "Comparator for the cycle start date."
required = false

[[inputs]]
name = "team"
type = "TeamFilter"
description = "Filters that the cycles team must satisfy."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
