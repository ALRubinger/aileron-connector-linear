+++
name = "favorite-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/favorite-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["favorite_update"]

[match]
intent = "Updates a favorite's position, parent folder, or folder name."

[[execute]]
id = "favorite"
connector = "github://ALRubinger/aileron-connector-linear"
op = "favorite_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "folderName"
type = "String"
description = "The name of the favorite folder."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the favorite to update."
required = true

[[inputs]]
name = "parentId"
type = "String"
description = "The identifier (in UUID v4 format) of the folder to move the favorite under."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The position of the item in the favorites list."
required = false
+++
