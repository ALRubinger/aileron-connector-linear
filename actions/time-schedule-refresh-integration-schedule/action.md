+++
name = "time-schedule-refresh-integration-schedule"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/time-schedule-refresh-integration-schedule@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["time_schedule_refresh_integration_schedule"]

[match]
intent = "Refresh the integration schedule information."

[[execute]]
id = "time"
connector = "github://ALRubinger/aileron-connector-linear"
op = "time_schedule_refresh_integration_schedule"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the time schedule to refresh."
required = true
+++
