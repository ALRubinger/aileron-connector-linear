+++
name = "oauth-application-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/oauth-application-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["oauth_application_create"]

[match]
intent = "[ALPHA] Creates an OAuth application owned by the calling OAuth application."

[[execute]]
id = "oauth"
connector = "github://ALRubinger/aileron-connector-linear"
op = "oauth_application_create"
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
required = true

[[inputs]]
name = "developerUrl"
type = "String"
description = "URL of the developer's website, homepage, or documentation."
required = true

[[inputs]]
name = "idempotencyKey"
type = "String"
description = "Optional client-supplied idempotency key. Reusing the same key with the same managing OAuth application returns the existing OAuth application instead of creating a duplicate. The key does not apply to archived applications."
required = false

[[inputs]]
name = "imageUrl"
type = "String"
description = "URL of the OAuth application's icon."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The OAuth application's name."
required = true

[[inputs]]
name = "redirectUris"
type = "String"
description = "Allowed redirect URIs for OAuth authorization flows."
required = true

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
