+++
name = "attachment-link-slack"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-link-slack@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_link_slack"]

[match]
intent = "Link an existing Slack message to an issue. This creates a rich attachment using the workspace's Slack integration."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_link_slack"
idempotent = false

[approval]
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
description = "The issue to which to link the Slack message. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = true

[[inputs]]
name = "syncToCommentThread"
type = "Boolean"
description = "Whether to begin syncing the message's Slack thread with a comment thread on the issue."
required = false

[[inputs]]
name = "title"
type = "String"
description = "The title to use for the attachment."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The Slack message URL for the message to link."
required = true
+++
