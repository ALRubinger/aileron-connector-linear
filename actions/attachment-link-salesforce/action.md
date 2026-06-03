+++
name = "attachment-link-salesforce"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-link-salesforce@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_link_salesforce"]

[match]
intent = "Link an existing Salesforce case to an issue. This creates a rich attachment using the workspace's Salesforce integration."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_link_salesforce"
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
description = "The issue for which to link the Salesforce case. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = true

[[inputs]]
name = "title"
type = "String"
description = "The title to use for the attachment."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The URL of the Salesforce case to link."
required = true
+++
