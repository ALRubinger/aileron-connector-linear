+++
name = "release-search"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-search@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_search"]

[match]
intent = "Search releases with optional text matching against name, version, and pipeline name. When no search term is provided, returns releases ordered by stage priority (started > planned > completed > canceled)."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_search"
idempotent = true

[[inputs]]
name = "and"
type = "ReleaseFilter"
description = "Compound filters, all of which need to be matched by the release."
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
description = "Maximum results. Capped at 50."
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
name = "term"
type = "String"
description = "Search term to match against release name, version, and pipeline name. When omitted, returns releases ordered by stage priority."
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
