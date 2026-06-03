+++
name = "organization"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization"]

[match]
intent = "The authenticated user's workspace."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization"
idempotent = true
+++
