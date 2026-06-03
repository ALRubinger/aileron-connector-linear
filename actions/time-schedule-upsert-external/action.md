+++
name = "time-schedule-upsert-external"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/time-schedule-upsert-external@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["time_schedule_upsert_external"]

[match]
intent = "Upsert an external time schedule."

[[execute]]
id = "time"
connector = "github://ALRubinger/aileron-connector-linear"
op = "time_schedule_upsert_external"
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
name = "name"
type = "String"
description = "The name of the schedule."
required = false
+++
