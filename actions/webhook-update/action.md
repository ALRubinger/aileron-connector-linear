+++
name = "webhook-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/webhook-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["webhook_update"]

[match]
intent = "Updates an existing Webhook."

[[execute]]
id = "webhook"
connector = "github://ALRubinger/aileron-connector-linear"
op = "webhook_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "enabled"
type = "Boolean"
description = "Whether this webhook is enabled."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the Webhook."
required = true

[[inputs]]
name = "label"
type = "String"
description = "Label for the webhook."
required = false

[[inputs]]
name = "resourceTypes"
type = "String"
description = "List of resources the webhook should subscribe to."
required = false

[[inputs]]
name = "secret"
type = "String"
description = "A secret token used to sign the webhook payload."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The URL that will be called on data changes."
required = false
+++
