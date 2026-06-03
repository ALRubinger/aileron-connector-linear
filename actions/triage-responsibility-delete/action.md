+++
name = "triage-responsibility-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/triage-responsibility-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["triage_responsibility_delete"]

[match]
intent = "Deletes a triage responsibility."

[[execute]]
id = "triage"
connector = "github://ALRubinger/aileron-connector-linear"
op = "triage_responsibility_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the triage responsibility to delete."
required = true
+++
