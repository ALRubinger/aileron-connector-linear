+++
name = "external-user"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/external-user@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["external_user"]

[match]
intent = "Retrieves a single external user by their identifier."

[[execute]]
id = "external"
connector = "github://ALRubinger/aileron-connector-linear"
op = "external_user"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the external user to retrieve."
required = true
+++
