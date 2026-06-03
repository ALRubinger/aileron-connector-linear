+++
name = "release-note"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-note@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_note"]

[match]
intent = "Fetch a release note by its UUID or slug identifier."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_note"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
