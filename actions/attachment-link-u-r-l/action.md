+++
name = "attachment-link-u-r-l"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-link-u-r-l@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_link_u_r_l"]

[match]
intent = "Link any URL to an issue. If the workspace has a matching integration configured and the URL is recognized (e.g., Zendesk, GitHub, Slack), a rich attachment will be created that enables features like automated status updates. Otherwise, a basic attachment is created."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_link_u_r_l"
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
description = "The id for the attachment."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The issue for which to link the url. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = true

[[inputs]]
name = "title"
type = "String"
description = "The title to use for the attachment."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The url to link."
required = true
+++
