+++
name = "team-unarchive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team-unarchive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team_unarchive"]

[match]
intent = "Unarchives a team and cancels deletion."

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team_unarchive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the team to unarchive."
required = true
+++
