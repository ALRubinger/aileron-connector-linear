+++
name = "integration-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_archive"]

[match]
intent = "Archives an integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the integration to archive."
required = true
+++
