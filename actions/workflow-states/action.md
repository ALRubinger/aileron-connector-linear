+++
name = "workflow-states"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/workflow-states@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["workflow_states"]

[match]
intent = "All issue workflow states (issue statuses). Returns a paginated list of workflow states visible to the authenticated user, across all teams they have access to."

[[execute]]
id = "workflow"
connector = "github://ALRubinger/aileron-connector-linear"
op = "workflow_states"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "WorkflowStateFilter"
description = "Compound filters, all of which need to be matched by the workflow state."
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
name = "description"
type = "StringComparator"
description = "Comparator for the workflow state description."
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
name = "issues"
type = "IssueCollectionFilter"
description = "Filters that the workflow states issues must satisfy."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "name"
type = "StringComparator"
description = "Comparator for the workflow state name."
required = false

[[inputs]]
name = "or"
type = "WorkflowStateFilter"
description = "Compound filters, one of which need to be matched by the workflow state."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "position"
type = "NumberComparator"
description = "Comparator for the workflow state position."
required = false

[[inputs]]
name = "team"
type = "TeamFilter"
description = "Filters that the workflow states team must satisfy."
required = false

[[inputs]]
name = "type"
type = "StringComparator"
description = "Comparator for the workflow state type. Possible values are \"triage\", \"backlog\", \"unstarted\", \"started\", \"completed\", \"canceled\", \"duplicate\"."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
