+++
name = "customer-needs"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-needs@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_needs"]

[match]
intent = "All customer needs in the workspace, with optional filtering."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_needs"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "CustomerNeedFilter"
description = "Compound filters, all of which need to be matched by the customer need."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "comment"
type = "NullableCommentFilter"
description = "Filters that the need's comment must satisfy."
required = false

[[inputs]]
name = "createdAt"
type = "DateComparator"
description = "Comparator for the created at date."
required = false

[[inputs]]
name = "customer"
type = "NullableCustomerFilter"
description = "Filters that the need's customer must satisfy."
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
name = "issue"
type = "NullableIssueFilter"
description = "Filters that the need's issue must satisfy."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "or"
type = "CustomerNeedFilter"
description = "Compound filters, one of which need to be matched by the customer need."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "priority"
type = "NumberComparator"
description = "Comparator for the customer need priority."
required = false

[[inputs]]
name = "project"
type = "NullableProjectFilter"
description = "Filters that the need's project must satisfy."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
