+++
name = "cycle-start-upcoming-cycle-today"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/cycle-start-upcoming-cycle-today@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["cycle_start_upcoming_cycle_today"]

[match]
intent = "Starts the upcoming cycle as of midnight today. Completes the previous cycle if it has not yet ended. Only the next upcoming (not yet started) cycle for the team can be started."

[[execute]]
id = "cycle"
connector = "github://ALRubinger/aileron-connector-linear"
op = "cycle_start_upcoming_cycle_today"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the cycle to start as of midnight today. Must be the upcoming cycle."
required = true
+++
