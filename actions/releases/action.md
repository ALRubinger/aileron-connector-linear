+++
name = "releases"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/releases@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["releases"]

[match]
intent = "All releases in the workspace, with optional filtering and sorting."

[[execute]]
id = "releases"
connector = "github://ALRubinger/aileron-connector-linear"
op = "releases"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "ReleaseFilter"
description = "Compound filters, all of which need to be matched by the release."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "completedAt"
type = "NullableDateComparator"
description = "Comparator for the release completion date."
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
name = "hasReleaseNotes"
type = "BooleanComparator"
description = "Comparator for whether the release is covered by any (non-archived) release note. Filter with `{ eq: false }` to retrieve releases that do not yet have release notes attached."
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
description = "Comparator for the release name."
required = false

[[inputs]]
name = "or"
type = "ReleaseFilter"
description = "Compound filters, one of which need to be matched by the release."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "pipeline"
type = "ReleasePipelineFilter"
description = "Filters that the release's pipeline must satisfy."
required = false

[[inputs]]
name = "stage"
type = "ReleaseStageFilter"
description = "Filters that the release's stage must satisfy."
required = false

[[inputs]]
name = "stage"
type = "ReleaseStageSort"
description = "Sort by release stage"
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false

[[inputs]]
name = "version"
type = "StringComparator"
description = "Comparator for the release version."
required = false
+++
