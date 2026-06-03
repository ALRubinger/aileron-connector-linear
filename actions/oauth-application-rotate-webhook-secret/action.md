+++
name = "oauth-application-rotate-webhook-secret"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/oauth-application-rotate-webhook-secret@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["oauth_application_rotate_webhook_secret"]

[match]
intent = "[ALPHA] Rotates the webhook signing secret for an OAuth application created by the calling OAuth application."

[[execute]]
id = "oauth"
connector = "github://ALRubinger/aileron-connector-linear"
op = "oauth_application_rotate_webhook_secret"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the OAuth application for which to rotate the webhook secret."
required = true
+++
