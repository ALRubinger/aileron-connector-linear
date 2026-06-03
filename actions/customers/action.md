+++
name = "customers"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customers@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customers"]

[match]
intent = "All customers in the workspace, with optional filtering and sorting."

[[execute]]
id = "customers"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customers"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "CustomerFilter"
description = "Compound filters, all of which need to be matched by the customer."
required = false

[[inputs]]
name = "approximateNeedCount"
type = "ApproximateNeedCountSort"
description = "Sort by approximate customer need count"
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
name = "createdAt"
type = "CustomerCreatedAtSort"
description = "Sort by customer creation date"
required = false

[[inputs]]
name = "domains"
type = "StringArrayComparator"
description = "Comparator for the customer's domains."
required = false

[[inputs]]
name = "externalIds"
type = "StringArrayComparator"
description = "Comparator for the customer's external IDs."
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
type = "StringComparator"
description = "Comparator for the customer name."
required = false

[[inputs]]
name = "name"
type = "NameSort"
description = "Sort by name"
required = false

[[inputs]]
name = "needs"
type = "CustomerNeedCollectionFilter"
description = "Filters that the customer's needs must satisfy."
required = false

[[inputs]]
name = "or"
type = "CustomerFilter"
description = "Compound filters, one of which need to be matched by the customer."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "owner"
type = "NullableUserFilter"
description = "Filters that the customer owner must satisfy."
required = false

[[inputs]]
name = "owner"
type = "OwnerSort"
description = "Sort by owner name"
required = false

[[inputs]]
name = "revenue"
type = "RevenueSort"
description = "Sort by customer generated revenue"
required = false

[[inputs]]
name = "revenue"
type = "NumberComparator"
description = "Comparator for the customer generated revenue."
required = false

[[inputs]]
name = "size"
type = "NumberComparator"
description = "Comparator for the customer size."
required = false

[[inputs]]
name = "size"
type = "SizeSort"
description = "Sort by customer size"
required = false

[[inputs]]
name = "slackChannelId"
type = "StringComparator"
description = "Comparator for the customer slack channel ID."
required = false

[[inputs]]
name = "status"
type = "CustomerStatusFilter"
description = "Filters that the customer's status must satisfy."
required = false

[[inputs]]
name = "status"
type = "CustomerStatusSort"
description = "Sort by customer status"
required = false

[[inputs]]
name = "tier"
type = "CustomerTierFilter"
description = "Filters that the customer's tier must satisfy."
required = false

[[inputs]]
name = "tier"
type = "TierSort"
description = "Sort by customer tier"
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
