+++
name = "semantic-search"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/semantic-search@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["semantic_search"]

[match]
intent = "Search for issues, projects, initiatives, and documents using natural language. Uses vector-based semantic search with optional full-text search and reranking. Results can be filtered by type and by entity-specific filters. Rate-limited to 30 requests per minute."

[[execute]]
id = "semantic"
connector = "github://ALRubinger/aileron-connector-linear"
op = "semantic_search"
idempotent = true

[[inputs]]
name = "documents"
type = "DocumentFilter"
description = "Filters applied to documents."
required = false

[[inputs]]
name = "includeArchived"
type = "Boolean"
description = "Whether to include archived results in the search (default: false)."
required = false

[[inputs]]
name = "initiatives"
type = "InitiativeFilter"
description = "Filters applied to initiatives."
required = false

[[inputs]]
name = "issues"
type = "IssueFilter"
description = "Filters applied to issues."
required = false

[[inputs]]
name = "maxResults"
type = "Int"
description = "The maximum number of results to return (default: 50)."
required = false

[[inputs]]
name = "projects"
type = "ProjectFilter"
description = "Filters applied to projects."
required = false

[[inputs]]
name = "query"
type = "String"
description = "Search query to look for."
required = true

[[inputs]]
name = "types"
type = "SemanticSearchResultType"
description = "The types of results to return (default: all)."
required = false
+++
