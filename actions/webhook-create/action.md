+++
name = "webhook-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/webhook-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["webhook_create"]

[match]
intent = "Creates a new webhook subscription for the workspace. Requires specifying a URL, resource types to subscribe to, and either a specific team or all public teams."

[[execute]]
id = "webhook"
connector = "github://ALRubinger/aileron-connector-linear"
op = "webhook_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "allPublicTeams"
type = "Boolean"
description = "Whether this webhook is enabled for all public teams."
required = false

[[inputs]]
name = "enabled"
type = "Boolean"
description = "Whether this webhook is enabled."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "label"
type = "String"
description = "Label for the webhook."
required = false

[[inputs]]
name = "resourceTypes"
type = "String"
description = "List of resources the webhook should subscribe to."
required = true

[[inputs]]
name = "secret"
type = "String"
description = "A secret token used to sign the webhook payload."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier or key of the team associated with the Webhook."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The URL that will be called on data changes."
required = true
+++
