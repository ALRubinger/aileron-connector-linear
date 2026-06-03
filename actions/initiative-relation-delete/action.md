+++
name = "initiative-relation-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-relation-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_relation_delete"]

[match]
intent = "Deletes an initiative relation."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_relation_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the initiative relation to delete."
required = true
+++
