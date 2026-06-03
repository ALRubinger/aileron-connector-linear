+++
name = "organization-start-trial"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-start-trial@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_start_trial"]

[match]
intent = "[DEPRECATED] Starts a trial for the workspace."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_start_trial"
idempotent = false

[approval]
required = true
+++
