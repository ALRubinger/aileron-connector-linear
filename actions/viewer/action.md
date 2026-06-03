+++
name = "viewer"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/viewer@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["get_viewer"]

[match]
intent = "look up the Linear user currently authenticated to the connector"

[[execute]]
id = "viewer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "viewer"
idempotent = true
+++
