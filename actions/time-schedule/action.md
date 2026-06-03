+++
name = "time-schedule"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/time-schedule@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["time_schedule"]

[match]
intent = "A specific time schedule."

[[execute]]
id = "time"
connector = "github://ALRubinger/aileron-connector-linear"
op = "time_schedule"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the time schedule to retrieve."
required = true
+++
