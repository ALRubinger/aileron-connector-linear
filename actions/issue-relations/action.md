+++
name = "issue-relations"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-relations@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_relations"]

[match]
intent = "All issue relations. Returns a paginated list of all issue relations (blocks, blocked by, relates to, duplicates) visible to the authenticated user."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_relations"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "first"
type = "Int"
description = "The number of items to forward paginate (used with after). Defaults to 50."
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
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false
+++
