+++
name = "templates-for-integration"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/templates-for-integration@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["templates_for_integration"]

[match]
intent = "Returns all templates that are associated with the integration type."

[[execute]]
id = "templates"
connector = "github://ALRubinger/aileron-connector-linear"
op = "templates_for_integration"
idempotent = true

[[inputs]]
name = "integrationType"
type = "String"
description = "The type of integration for which to return associated templates."
required = true
+++
