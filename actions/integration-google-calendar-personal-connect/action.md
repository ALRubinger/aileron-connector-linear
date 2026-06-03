+++
name = "integration-google-calendar-personal-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-google-calendar-personal-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_google_calendar_personal_connect"]

[match]
intent = "[Internal] Connects the Google Calendar to the user to this Linear account via OAuth2."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_google_calendar_personal_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "[Internal] The Google OAuth code."
required = true
+++
