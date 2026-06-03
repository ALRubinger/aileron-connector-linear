+++
name = "archived-teams"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/archived-teams@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["archived_teams"]

[match]
intent = "[Internal] All archived teams of the workspace."

[[execute]]
id = "archived"
connector = "github://ALRubinger/aileron-connector-linear"
op = "archived_teams"
idempotent = true
+++
