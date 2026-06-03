+++
name = "users"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/users@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["users"]

[match]
intent = "All users in the workspace. Supports filtering, sorting, and pagination."

[[execute]]
id = "users"
connector = "github://ALRubinger/aileron-connector-linear"
op = "users"
idempotent = true

[[inputs]]
name = "active"
type = "BooleanComparator"
description = "Comparator for the user's activity status."
required = false

[[inputs]]
name = "admin"
type = "BooleanComparator"
description = "Comparator for the user's admin status."
required = false

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "UserFilter"
description = "Compound filters, all of which need to be matched by the user."
required = false

[[inputs]]
name = "app"
type = "BooleanComparator"
description = "Comparator for the user's app status."
required = false

[[inputs]]
name = "assignedIssues"
type = "IssueCollectionFilter"
description = "Filters that the users assigned issues must satisfy."
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
name = "displayName"
type = "StringComparator"
description = "Comparator for the user's display name."
required = false

[[inputs]]
name = "displayName"
type = "UserDisplayNameSort"
description = "Sort by user display name"
required = false

[[inputs]]
name = "email"
type = "StringComparator"
description = "Comparator for the user's email."
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
name = "includeDisabled"
type = "Boolean"
description = "Should query return disabled/suspended users (default: false)."
required = false

[[inputs]]
name = "invited"
type = "BooleanComparator"
description = "Comparator for the user's invited status."
required = false

[[inputs]]
name = "isInvited"
type = "BooleanComparator"
description = "Comparator for the user's invited status."
required = false

[[inputs]]
name = "isMe"
type = "BooleanComparator"
description = "Filter based on the currently authenticated user. Set to true to filter for the authenticated user, false for any other user."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "name"
type = "StringComparator"
description = "Comparator for the user's name."
required = false

[[inputs]]
name = "name"
type = "UserNameSort"
description = "Sort by user name"
required = false

[[inputs]]
name = "or"
type = "UserFilter"
description = "Compound filters, one of which need to be matched by the user."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "owner"
type = "BooleanComparator"
description = "Comparator for the user's owner status."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
