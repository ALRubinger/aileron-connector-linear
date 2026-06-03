+++
name = "organization-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_updated"]

[match]
intent = "Triggered when an organization is updated"

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_updated"
idempotent = true
+++
