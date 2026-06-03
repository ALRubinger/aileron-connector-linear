+++
name = "integration-launch-darkly-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-launch-darkly-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_launch_darkly_connect"]

[match]
intent = "[INTERNAL] Integrates the workspace with LaunchDarkly."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_launch_darkly_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The LaunchDarkly OAuth code."
required = true

[[inputs]]
name = "environment"
type = "String"
description = "The LaunchDarkly environment."
required = true

[[inputs]]
name = "projectKey"
type = "String"
description = "The LaunchDarkly project key."
required = true
+++
