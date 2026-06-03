+++
name = "release-note-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-note-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_note_delete"]

[match]
intent = "Deletes a release note."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_note_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the release note to delete."
required = true
+++
