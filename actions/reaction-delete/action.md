+++
name = "reaction-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/reaction-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["reaction_delete"]

[match]
intent = "Deletes a reaction."

[[execute]]
id = "reaction"
connector = "github://ALRubinger/aileron-connector-linear"
op = "reaction_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the reaction to delete."
required = true
+++
