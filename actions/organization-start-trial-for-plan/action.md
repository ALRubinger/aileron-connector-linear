+++
name = "organization-start-trial-for-plan"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-start-trial-for-plan@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_start_trial_for_plan"]

[match]
intent = "Starts a trial for the workspace on the specified plan type. The workspace must not already be on a paid plan or in an active trial."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_start_trial_for_plan"
idempotent = false

[approval]
required = true

[[inputs]]
name = "planType"
type = "String"
description = "The plan type to trial."
required = true
+++
