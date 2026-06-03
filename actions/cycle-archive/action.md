+++
name = "cycle-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/cycle-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["cycle_archive"]

[match]
intent = "Archives a cycle. All issues currently assigned to the cycle are unlinked from it before archiving."

[[execute]]
id = "cycle"
connector = "github://ALRubinger/aileron-connector-linear"
op = "cycle_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the cycle to archive."
required = true
+++
