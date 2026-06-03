+++
name = "integration-github-import-refresh"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-github-import-refresh@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_github_import_refresh"]

[match]
intent = "Refreshes the data for a GitHub import integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_github_import_refresh"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The id of the integration to update."
required = true
+++
