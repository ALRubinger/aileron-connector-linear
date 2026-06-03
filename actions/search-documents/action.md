+++
name = "search-documents"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/search-documents@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["search_documents"]

[match]
intent = "Search documents by text query using full-text and vector search. Results are ranked by relevance unless an orderBy parameter is specified. Rate-limited to 30 requests per minute."

[[execute]]
id = "search"
connector = "github://ALRubinger/aileron-connector-linear"
op = "search_documents"
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
name = "includeComments"
type = "Boolean"
description = "Should associated comments be searched (default: false)."
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

[[inputs]]
name = "teamId"
type = "String"
description = "UUID of a team to boost in search results. Results from this team are ranked higher. If null, no team boosting is applied."
required = false

[[inputs]]
name = "term"
type = "String"
description = "Search string to look for."
required = true
+++
