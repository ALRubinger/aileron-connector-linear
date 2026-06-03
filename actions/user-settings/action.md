+++
name = "user-settings"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-settings@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_settings"]

[match]
intent = "The authenticated user's notification and UI settings."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_settings"
idempotent = true
+++
