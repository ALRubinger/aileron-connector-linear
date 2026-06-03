+++
name = "release-note-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-note-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_note_update"]

[match]
intent = "Updates a release note."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_note_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "content"
type = "String"
description = "The release note body as markdown."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the release note to update."
required = true

[[inputs]]
name = "rangeFromReleaseId"
type = "String"
description = "Oldest release (by createdAt) of the new range. Paired with rangeToReleaseId."
required = false

[[inputs]]
name = "rangeToReleaseId"
type = "String"
description = "Newest release (by createdAt) of the new range. Paired with rangeFromReleaseId."
required = false

[[inputs]]
name = "releaseIds"
type = "String"
description = "Explicit release IDs to set. Mutually exclusive with rangeFromReleaseId/rangeToReleaseId."
required = false

[[inputs]]
name = "title"
type = "String"
description = "Optional user-supplied title."
required = false
+++
