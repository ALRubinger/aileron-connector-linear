+++
name = "fetch-data"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/fetch-data@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["fetch_data"]

[match]
intent = "[Internal] Fetch an arbitrary set of data using natural language query. Be specific about what you want including properties for each entity, sort order, filters, limit and properties."

[[execute]]
id = "fetch"
connector = "github://ALRubinger/aileron-connector-linear"
op = "fetch_data"
idempotent = true

[[inputs]]
name = "query"
type = "String"
description = "Natural language query describing what data to fetch.\n\n    Examples:\n    - \"All issues for the project with id 12345678-1234-1234-1234-123456789abc including comments\"\n    - \"The latest project update for each project that's a part of the initiative with id 12345678-1234-1234-1234-123456789abc, including it's sub-initiatives\""
required = true
+++
