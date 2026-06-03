+++
name = "integration-salesforce-metadata-refresh"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-salesforce-metadata-refresh@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_salesforce_metadata_refresh"]

[match]
intent = "[INTERNAL] Refreshes the Salesforce integration metadata."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_salesforce_metadata_refresh"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The ID of the integration to refresh metadata for."
required = true
+++
