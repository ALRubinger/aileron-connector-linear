+++
name = "integration-google-sheets"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-google-sheets@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_google_sheets"]

[match]
intent = "Integrates the workspace with Google Sheets."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_google_sheets"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Google OAuth code."
required = true
+++
