+++
name = "organization-delete-challenge"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-delete-challenge@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_delete_challenge"]

[match]
intent = "Requests a deletion confirmation code for the workspace. The code is sent to the requesting user's email and must be used with the organizationDelete mutation to confirm deletion."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_delete_challenge"
idempotent = false

[approval]
required = true
+++
