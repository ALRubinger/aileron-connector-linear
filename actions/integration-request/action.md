+++
name = "integration-request"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-request@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_request"]

[match]
intent = "Requests a currently unavailable integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_request"
idempotent = false

[approval]
required = true

[[inputs]]
name = "email"
type = "String"
description = "Email associated with the request."
required = false

[[inputs]]
name = "name"
type = "String"
description = "Name of the requested integration."
required = true
+++
