+++
name = "team"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/team@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["team"]

[match]
intent = "Fetches a specific team by its ID."

[[execute]]
id = "team"
connector = "github://ALRubinger/aileron-connector-linear"
op = "team"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the team to retrieve."
required = true
+++
