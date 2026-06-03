+++
name = "triage-responsibility"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/triage-responsibility@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["triage_responsibility"]

[match]
intent = "A specific triage responsibility."

[[execute]]
id = "triage"
connector = "github://ALRubinger/aileron-connector-linear"
op = "triage_responsibility"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the triage responsibility to retrieve."
required = true
+++
