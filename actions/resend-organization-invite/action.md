+++
name = "resend-organization-invite"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/resend-organization-invite@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["resend_organization_invite"]

[match]
intent = "Re-sends a workspace invitation email for the specified invite ID."

[[execute]]
id = "resend"
connector = "github://ALRubinger/aileron-connector-linear"
op = "resend_organization_invite"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the organization invite to re-send."
required = true
+++
