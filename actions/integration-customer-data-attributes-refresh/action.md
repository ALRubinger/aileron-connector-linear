+++
name = "integration-customer-data-attributes-refresh"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-customer-data-attributes-refresh@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_customer_data_attributes_refresh"]

[match]
intent = "[INTERNAL] Refreshes the customer data attributes from the specified integration service."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_customer_data_attributes_refresh"
idempotent = false

[approval]
required = true

[[inputs]]
name = "service"
type = "String"
description = "The integration service to refresh customer data attributes from."
required = true
+++
