+++
name = "organization-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_delete"]

[match]
intent = "Permanently deletes the workspace and all its data. Requires a valid deletion code obtained from organizationDeleteChallenge. This action is irreversible."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "deletionCode"
type = "String"
description = "The deletion code to confirm operation."
required = true
+++
