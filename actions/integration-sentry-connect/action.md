+++
name = "integration-sentry-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-sentry-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_sentry_connect"]

[match]
intent = "Integrates the workspace with Sentry."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_sentry_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Sentry grant code that's exchanged for OAuth tokens."
required = true

[[inputs]]
name = "installationId"
type = "String"
description = "The Sentry installationId to connect with."
required = true

[[inputs]]
name = "organizationSlug"
type = "String"
description = "The slug of the Sentry organization being connected."
required = true
+++
