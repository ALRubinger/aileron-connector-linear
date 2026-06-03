+++
name = "webhook"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/webhook@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["webhook"]

[match]
intent = "Retrieves a single webhook by its identifier."

[[execute]]
id = "webhook"
connector = "github://ALRubinger/aileron-connector-linear"
op = "webhook"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the webhook to retrieve."
required = true
+++
