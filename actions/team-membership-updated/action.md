+++
name = "team-membership-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-membership-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_membership_updated"]

[match]
intent = "Triggered when a team membership is updated"

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_membership_updated"
idempotent = true
+++
