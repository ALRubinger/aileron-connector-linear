+++
name = "audit-entries"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/audit-entries@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["audit_entries"]

[match]
intent = "All audit log entries."

[[execute]]
id = "audit"
connector = "github://ALRubinger/aileron-connector-linear"
op = "audit_entries"
idempotent = true

[[inputs]]
name = "actor"
type = "NullableUserFilter"
description = "Filters that the audit entry actor must satisfy."
required = false

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "AuditEntryFilter"
description = "Compound filters, all of which need to be matched by the issue."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "countryCode"
type = "StringComparator"
description = "Comparator for the country code."
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
name = "ip"
type = "StringComparator"
description = "Comparator for the IP address."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "or"
type = "AuditEntryFilter"
description = "Compound filters, one of which need to be matched by the issue."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "type"
type = "StringComparator"
description = "Comparator for the type."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
