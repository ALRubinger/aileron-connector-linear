+++
name = "leave-organization"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/leave-organization@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["leave_organization"]

[match]
intent = "Leave a workspace."

[[execute]]
id = "leave"
connector = "github://ALRubinger/aileron-connector-linear"
op = "leave_organization"
idempotent = false

[approval]
required = true

[[inputs]]
name = "organizationId"
type = "String"
description = "ID of the workspace to leave."
required = true
+++
