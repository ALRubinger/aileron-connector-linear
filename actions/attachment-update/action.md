+++
name = "attachment-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_update"]

[match]
intent = "Updates an existing issue attachment."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "iconUrl"
type = "String"
description = "An icon url to display with the attachment. Should be of jpg or png format. Maximum of 1MB in size. Dimensions should be 20x20px for optimal display quality."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the attachment to update."
required = true

[[inputs]]
name = "metadata"
type = "JSONObject"
description = "Attachment metadata object with string and number values."
required = false

[[inputs]]
name = "subtitle"
type = "String"
description = "The attachment subtitle."
required = false

[[inputs]]
name = "title"
type = "String"
description = "The attachment title."
required = true
+++
