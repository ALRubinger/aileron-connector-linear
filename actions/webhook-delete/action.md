+++
name = "webhook-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/webhook-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["webhook_delete"]

[match]
intent = "Deletes a Webhook."

[[execute]]
id = "webhook"
connector = "github://ALRubinger/aileron-connector-linear"
op = "webhook_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the Webhook to delete."
required = true
+++
