+++
name = "create-organization-from-onboarding"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/create-organization-from-onboarding@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["create_organization_from_onboarding"]

[match]
intent = "Creates a new workspace from onboarding."

[[execute]]
id = "create"
connector = "github://ALRubinger/aileron-connector-linear"
op = "create_organization_from_onboarding"
idempotent = false

[approval]
required = true

[[inputs]]
name = "companyRole"
type = "String"
description = ""
required = false

[[inputs]]
name = "companySize"
type = "String"
description = ""
required = false

[[inputs]]
name = "domainAccess"
type = "Boolean"
description = "Whether the organization should allow email domain access."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the organization."
required = true

[[inputs]]
name = "sessionId"
type = "String"
description = "PostHog session ID for attribution tracking."
required = false

[[inputs]]
name = "timezone"
type = "String"
description = "The timezone of the organization, passed in by client."
required = false

[[inputs]]
name = "urlKey"
type = "String"
description = "The URL key of the organization."
required = true

[[inputs]]
name = "utm"
type = "String"
description = "JSON serialized UTM parameters associated with the creation of the workspace."
required = false

[[inputs]]
name = "utmFirstTouch"
type = "String"
description = "JSON serialized UTM parameters captured on the user's first visit to the marketing site (first-touch attribution)."
required = false
+++
