+++
name = "issue-labels"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-labels@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_labels"]

[match]
intent = "All issue labels. Returns a paginated list of labels visible to the authenticated user, including both workspace-level and team-scoped labels."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_labels"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "IssueLabelFilter"
description = "Compound filters, all of which need to be matched by the label."
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
name = "creator"
type = "NullableUserFilter"
description = "Filters that the issue labels creator must satisfy."
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
name = "isGroup"
type = "BooleanComparator"
description = "Comparator for whether the label is a group label."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "name"
type = "StringComparator"
description = "Comparator for the name."
required = false

[[inputs]]
name = "or"
type = "IssueLabelFilter"
description = "Compound filters, one of which need to be matched by the label."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "parent"
type = "IssueLabelFilter"
description = "Filters that the issue label's parent label must satisfy."
required = false

[[inputs]]
name = "team"
type = "NullableTeamFilter"
description = "Filters that the issue labels team must satisfy."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
