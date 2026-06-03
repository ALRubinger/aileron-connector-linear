+++
name = "integration-salesforce"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-salesforce@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_salesforce"]

[match]
intent = "Integrates the workspace with Salesforce."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_salesforce"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Salesforce OAuth code."
required = true

[[inputs]]
name = "codeVerifier"
type = "String"
description = "The PKCE code verifier matching the code_challenge sent on authorization."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Salesforce OAuth redirect URI."
required = true

[[inputs]]
name = "subdomain"
type = "String"
description = "The Salesforce installation subdomain."
required = true
+++
