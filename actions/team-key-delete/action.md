+++
name = "team-key-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-key-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_key_delete"]

[match]
intent = "Deletes a previously used team key. The active team key (the team's current key) cannot be deleted."

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_key_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the team key to delete."
required = true
+++
