+++
name = "resend-organization-invite-by-email"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/resend-organization-invite-by-email@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["resend_organization_invite_by_email"]

[match]
intent = "Re-sends a workspace invitation email to the specified email address, if an outstanding invite exists for that address."

[[execute]]
id = "resend"
connector = "github://ALRubinger/aileron-connector-linear"
op = "resend_organization_invite_by_email"
idempotent = false

[approval]
required = true

[[inputs]]
name = "email"
type = "String"
description = "The email address tied to the organization invite to re-send."
required = true
+++
