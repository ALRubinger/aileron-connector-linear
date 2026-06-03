+++
name = "team-cycles-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-cycles-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_cycles_delete"]

[match]
intent = "Deletes all cycle data for a team, disabling the cycles feature. This removes all cycles and their issue associations."

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_cycles_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the team, which cycles will be deleted."
required = true
+++
