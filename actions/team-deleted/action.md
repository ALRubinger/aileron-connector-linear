+++
name = "team-deleted"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-deleted@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_deleted"]

[match]
intent = "Triggered when a team is deleted"

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_deleted"
idempotent = true
+++
