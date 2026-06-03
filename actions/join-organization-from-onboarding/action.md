+++
name = "join-organization-from-onboarding"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/join-organization-from-onboarding@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["join_organization_from_onboarding"]

[match]
intent = "Join a workspace from onboarding."

[[execute]]
id = "join"
connector = "github://ALRubinger/aileron-connector-linear"
op = "join_organization_from_onboarding"
idempotent = false

[approval]
required = true

[[inputs]]
name = "inviteLink"
type = "String"
description = "An optional invite link for an organization."
required = false

[[inputs]]
name = "organizationId"
type = "String"
description = "The identifier of the organization."
required = true
+++
