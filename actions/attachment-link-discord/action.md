+++
name = "attachment-link-discord"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-link-discord@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_link_discord"]

[match]
intent = "Link an existing Discord message to an issue. This creates a rich attachment using the workspace's Discord integration."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_link_discord"
idempotent = false

[approval]
required = true

[[inputs]]
name = "channelId"
type = "String"
description = "The Discord channel ID for the message to link."
required = true

[[inputs]]
name = "createAsUser"
type = "String"
description = "Create attachment as a user with the provided name. This option is only available to OAuth applications creating attachments in `actor=app` mode."
required = false

[[inputs]]
name = "displayIconUrl"
type = "String"
description = "Provide an external user avatar URL. Can only be used in conjunction with the `createAsUser` options. This option is only available to OAuth applications creating comments in `actor=app` mode."
required = false

[[inputs]]
name = "id"
type = "String"
description = "Optional attachment ID that may be provided through the API."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The issue for which to link the Discord message. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = true

[[inputs]]
name = "messageId"
type = "String"
description = "The Discord message ID for the message to link."
required = true

[[inputs]]
name = "title"
type = "String"
description = "The title to use for the attachment."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The Discord message URL for the message to link."
required = true
+++
