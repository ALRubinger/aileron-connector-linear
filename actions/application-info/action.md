+++
name = "application-info"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/application-info@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["application_info"]

[match]
intent = "Retrieves public information about an OAuth application by its client ID. Used during the authorization flow to display application details to the user."

[[execute]]
id = "application"
connector = "github://ALRubinger/aileron-connector-linear"
op = "application_info"
idempotent = true

[[inputs]]
name = "clientId"
type = "String"
description = "The client ID of the application."
required = true
+++
