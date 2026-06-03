+++
name = "project-milestones"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-milestones@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_milestones"]

[match]
intent = "Returns all project milestones in the workspace, with optional filtering."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_milestones"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "ProjectMilestoneFilter"
description = "Compound filters, all of which need to be matched by the project milestone."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "createdAt"
type = "DateComparator"
description = "Comparator for the created at date."
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
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "name"
type = "NullableStringComparator"
description = "Comparator for the project milestone name."
required = false

[[inputs]]
name = "or"
type = "ProjectMilestoneFilter"
description = "Compound filters, one of which need to be matched by the project milestone."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "project"
type = "NullableProjectFilter"
description = "Filters that the project milestone's project must satisfy."
required = false

[[inputs]]
name = "targetDate"
type = "NullableDateComparator"
description = "Comparator for the project milestone target date."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
