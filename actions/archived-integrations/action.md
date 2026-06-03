+++
name = "archived-integrations"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/archived-integrations@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["archived_integrations"]

[match]
intent = "Recently archived GitHub integrations. Returns integrations archived within the last 30 days."

[[execute]]
id = "archived"
connector = "github://ALRubinger/aileron-connector-linear"
op = "archived_integrations"
idempotent = true
+++
