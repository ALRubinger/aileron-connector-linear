+++
name = "attachments"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachments@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachments"]

[match]
intent = "All issue attachments.\n\nTo get attachments for a given URL, use `attachmentsForURL` query."

[[execute]]
id = "attachments"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachments"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "AttachmentFilter"
description = "Compound filters, all of which need to be matched by the attachment."
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
description = "Filters that the attachments creator must satisfy."
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
name = "or"
type = "AttachmentFilter"
description = "Compound filters, one of which need to be matched by the attachment."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "sourceType"
type = "SourceTypeComparator"
description = "Comparator for the source type."
required = false

[[inputs]]
name = "subtitle"
type = "NullableStringComparator"
description = "Comparator for the subtitle."
required = false

[[inputs]]
name = "title"
type = "StringComparator"
description = "Comparator for the title."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false

[[inputs]]
name = "url"
type = "StringComparator"
description = "Comparator for the url."
required = false
+++
