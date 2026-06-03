+++
name = "sla-configurations"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/sla-configurations@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["sla_configurations"]

[match]
intent = "Active SLA configurations that can apply to the requested team."

[[execute]]
id = "sla"
connector = "github://ALRubinger/aileron-connector-linear"
op = "sla_configurations"
idempotent = true

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier or key of the team to evaluate SLA rules against."
required = true
+++
