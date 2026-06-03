+++
name = "attachment-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_create"]

[match]
intent = "Creates a new attachment, or updates existing if the same `url` and `issueId` is used. To create an integration-aware attachment, use the integration-specific mutations such as `attachmentLinkZendesk`, `attachmentLinkSlack`, or `attachmentLinkURL` instead."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "commentBody"
type = "String"
description = "Create a linked comment with markdown body."
required = false

[[inputs]]
name = "commentBodyData"
type = "JSONObject"
description = "[Internal] Create a linked comment with Prosemirror body. Please use `commentBody` instead."
required = false

[[inputs]]
name = "createAsUser"
type = "String"
description = "Create attachment as a user with the provided name. This option is only available to OAuth applications creating attachments in `actor=application` mode."
required = false

[[inputs]]
name = "displayIconUrl"
type = "String"
description = "Provide an external user avatar URL. Can only be used in conjunction with the `createAsUser` options. This option is only available to OAuth applications creating attachments in `actor=app` mode."
required = false

[[inputs]]
name = "groupBySource"
type = "Boolean"
description = "Indicates if attachments for the same source application should be grouped in the Linear UI."
required = false

[[inputs]]
name = "iconUrl"
type = "String"
description = "An icon url to display with the attachment. Should be of jpg or png format. Maximum of 1MB in size. Dimensions should be 20x20px for optimal display quality."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The issue to associate the attachment with. Can be a UUID or issue identifier (e.g., 'LIN-123')."
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

[[inputs]]
name = "url"
type = "String"
description = "Attachment location which is also used as an unique identifier for the attachment. If another attachment is created with the same `url` value, existing record is updated instead."
required = true
+++
