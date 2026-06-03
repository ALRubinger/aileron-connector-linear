+++
name = "failures-for-oauth-webhooks"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/failures-for-oauth-webhooks@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["failures_for_oauth_webhooks"]

[match]
intent = "[INTERNAL] Webhook failure events for webhooks that belong to an OAuth application. (last 50)"

[[execute]]
id = "failures"
connector = "github://ALRubinger/aileron-connector-linear"
op = "failures_for_oauth_webhooks"
idempotent = true

[[inputs]]
name = "oauthClientId"
type = "String"
description = "The identifier of the OAuth client to retrieve failures for."
required = true
+++
