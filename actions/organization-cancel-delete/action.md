+++
name = "organization-cancel-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-cancel-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_cancel_delete"]

[match]
intent = "Cancels a previously requested workspace deletion, if the workspace has not yet been fully deleted."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_cancel_delete"
idempotent = false

[approval]
required = true
+++
