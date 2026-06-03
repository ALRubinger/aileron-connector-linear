+++
name = "webhook-rotate-secret"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/webhook-rotate-secret@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["webhook_rotate_secret"]

[match]
intent = "Rotates the signing secret for a webhook, generating a new HMAC-SHA256 key and returning it. The old secret is immediately invalidated."

[[execute]]
id = "webhook"
connector = "github://ALRubinger/aileron-connector-linear"
op = "webhook_rotate_secret"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the Webhook to rotate the secret for."
required = true
+++
