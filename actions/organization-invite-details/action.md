+++
name = "organization-invite-details"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-invite-details@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_invite_details"]

[match]
intent = "Fetches public details of a specific workspace invite by its ID. This query does not require authentication and is used for invite acceptance flows."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_invite_details"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the invite to retrieve details for."
required = true
+++
