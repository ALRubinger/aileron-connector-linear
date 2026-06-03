+++
name = "team-membership"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-membership@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_membership"]

[match]
intent = "Fetches a specific team membership by its ID."

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_membership"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the team membership to retrieve."
required = true
+++
