+++
name = "entity-external-link-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/entity-external-link-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["entity_external_link_delete"]

[match]
intent = "Deletes an entity external link."

[[execute]]
id = "entity"
connector = "github://ALRubinger/aileron-connector-linear"
op = "entity_external_link_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the entity link to delete."
required = true
+++
