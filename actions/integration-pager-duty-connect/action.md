+++
name = "integration-pager-duty-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-pager-duty-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_pager_duty_connect"]

[match]
intent = "[INTERNAL] Integrates the workspace with PagerDuty."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_pager_duty_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The PagerDuty OAuth code."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The PagerDuty OAuth redirect URI."
required = true
+++
