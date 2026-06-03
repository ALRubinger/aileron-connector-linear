+++
name = "documents"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/documents@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["documents"]

[match]
intent = "All documents the user has access to in the workspace."

[[execute]]
id = "documents"
connector = "github://ALRubinger/aileron-connector-linear"
op = "documents"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "DocumentFilter"
description = "Compound filters, all of which need to be matched by the document."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "createdAt"
type = "DocumentCreatedAtSort"
description = "Sort by document creation date."
required = false

[[inputs]]
name = "createdAt"
type = "DateComparator"
description = "Comparator for the created at date."
required = false

[[inputs]]
name = "creator"
type = "DocumentCreatorSort"
description = "Sort by the document's creator (author)."
required = false

[[inputs]]
name = "creator"
type = "UserFilter"
description = "Filters that the document's creator must satisfy."
required = false

[[inputs]]
name = "cycle"
type = "CycleFilter"
description = "Filters that the document's cycle must satisfy."
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
name = "initiative"
type = "InitiativeFilter"
description = "Filters that the document's initiative must satisfy."
required = false

[[inputs]]
name = "issue"
type = "IssueFilter"
description = "Filters that the document's issue must satisfy."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "or"
type = "DocumentFilter"
description = "Compound filters, one of which need to be matched by the document."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "project"
type = "ProjectFilter"
description = "Filters that the document's project must satisfy."
required = false

[[inputs]]
name = "project"
type = "DocumentProjectSort"
description = "Sort by the document's parent project name."
required = false

[[inputs]]
name = "release"
type = "ReleaseFilter"
description = "Filters that the document's release must satisfy."
required = false

[[inputs]]
name = "slugId"
type = "StringComparator"
description = "Comparator for the document slug ID."
required = false

[[inputs]]
name = "team"
type = "NullableTeamFilter"
description = "Filters that the document's team must satisfy."
required = false

[[inputs]]
name = "title"
type = "StringComparator"
description = "Comparator for the document title."
required = false

[[inputs]]
name = "title"
type = "DocumentTitleSort"
description = "Sort by document title."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false

[[inputs]]
name = "updatedAt"
type = "DocumentUpdatedAtSort"
description = "Sort by document last-updated date."
required = false
+++
