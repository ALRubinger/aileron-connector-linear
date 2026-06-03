+++
name = "integration-intercom"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-intercom@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_intercom"]

[match]
intent = "Integrates the workspace with Intercom."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_intercom"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Intercom OAuth code."
required = true

[[inputs]]
name = "domainUrl"
type = "String"
description = "The Intercom domain URL to use for the integration. Defaults to app.intercom.com if not provided."
required = false

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Intercom OAuth redirect URI."
required = true
+++
