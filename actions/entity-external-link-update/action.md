+++
name = "entity-external-link-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/entity-external-link-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["entity_external_link_update"]

[match]
intent = "Updates an existing entity external link's URL, label, or sort order."

[[execute]]
id = "entity"
connector = "github://ALRubinger/aileron-connector-linear"
op = "entity_external_link_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the entity link to update."
required = true

[[inputs]]
name = "label"
type = "String"
description = "The label for the link."
required = false

[[inputs]]
name = "resourceFolderId"
type = "String"
description = "[Internal] The resource folder containing the link."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The order of the item in the entities resources list."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The URL of the link."
required = false
+++
