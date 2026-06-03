+++
name = "time-schedule-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/time-schedule-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["time_schedule_update"]

[match]
intent = "Updates a time schedule."

[[execute]]
id = "time"
connector = "github://ALRubinger/aileron-connector-linear"
op = "time_schedule_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "entries"
type = "TimeScheduleEntryInput"
description = "The schedule entries."
required = false

[[inputs]]
name = "externalId"
type = "String"
description = "The unique identifier of the external schedule."
required = false

[[inputs]]
name = "externalUrl"
type = "String"
description = "The URL to the external schedule."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the time schedule to update."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The name of the schedule."
required = false
+++
