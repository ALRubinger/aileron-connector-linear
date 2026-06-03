+++
name = "time-schedule-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/time-schedule-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["time_schedule_create"]

[match]
intent = "Creates a new time schedule."

[[execute]]
id = "time"
connector = "github://ALRubinger/aileron-connector-linear"
op = "time_schedule_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "entries"
type = "TimeScheduleEntryInput"
description = "The schedule entries."
required = true

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
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the schedule."
required = true
+++
