+++
name = "integration-pager-duty-refresh-schedule-mappings"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-pager-duty-refresh-schedule-mappings@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_pager_duty_refresh_schedule_mappings"]

[match]
intent = "[INTERNAL] Refresh PagerDuty schedule mappings."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_pager_duty_refresh_schedule_mappings"
idempotent = false

[approval]
required = true
+++
