+++
name = "oauth-application-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/oauth-application-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["oauth_application_update"]

[match]
intent = "[ALPHA] Updates an OAuth application created by the calling OAuth application."

[[execute]]
id = "oauth"
connector = "github://ALRubinger/aileron-connector-linear"
op = "oauth_application_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "description"
type = "String"
description = "User-facing description of the OAuth application."
required = false

[[inputs]]
name = "developer"
type = "String"
description = "Name of the developer or company that built the OAuth application."
required = false

[[inputs]]
name = "developerUrl"
type = "String"
description = "URL of the developer's website, homepage, or documentation."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the OAuth application to update."
required = true

[[inputs]]
name = "imageUrl"
type = "String"
description = "URL of the OAuth application's icon."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The OAuth application's name."
required = false

[[inputs]]
name = "redirectUris"
type = "String"
description = "Allowed redirect URIs for OAuth authorization flows."
required = false

[[inputs]]
name = "webhookEnabled"
type = "Boolean"
description = "Whether webhook delivery is enabled for this OAuth application."
required = false

[[inputs]]
name = "webhookResourceTypes"
type = "String"
description = "Resource types the OAuth application's webhooks subscribe to."
required = false

[[inputs]]
name = "webhookUrl"
type = "String"
description = "Webhook URL used for delivering webhook payloads."
required = false
+++
