+++
name = "organization-invite"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-invite@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_invite"]

[match]
intent = "Fetches a specific workspace invite by its ID."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_invite"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the invite to retrieve."
required = true
+++
