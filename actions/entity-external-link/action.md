+++
name = "entity-external-link"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/entity-external-link@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["entity_external_link"]

[match]
intent = "Retrieves a single entity external link by its identifier."

[[execute]]
id = "entity"
connector = "github://ALRubinger/aileron-connector-linear"
op = "entity_external_link"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the entity external link to retrieve."
required = true
+++
