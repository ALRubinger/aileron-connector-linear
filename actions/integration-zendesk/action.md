+++
name = "integration-zendesk"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-zendesk@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_zendesk"]

[match]
intent = "Integrates the workspace with Zendesk."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_zendesk"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Zendesk OAuth code."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Zendesk OAuth redirect URI."
required = true

[[inputs]]
name = "scope"
type = "String"
description = "The Zendesk OAuth scopes."
required = true

[[inputs]]
name = "subdomain"
type = "String"
description = "The Zendesk installation subdomain."
required = true
+++
