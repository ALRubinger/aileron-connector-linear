+++
name = "release-note-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-note-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_note_create"]

[match]
intent = "Creates a release note."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_note_create"
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
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "pipelineId"
type = "String"
description = "Identifier of the release pipeline."
required = true

[[inputs]]
name = "rangeFromReleaseId"
type = "String"
description = "Oldest release (by createdAt) to include. When paired with rangeToReleaseId, expands to every release in the pipeline within the createdAt window."
required = false

[[inputs]]
name = "rangeToReleaseId"
type = "String"
description = "Newest release (by createdAt) to include. Paired with rangeFromReleaseId."
required = false

[[inputs]]
name = "releaseIds"
type = "String"
description = "Explicit release IDs to include. Mutually exclusive with rangeFromReleaseId/rangeToReleaseId — exactly one of the two shapes must be provided."
required = false

[[inputs]]
name = "title"
type = "String"
description = "Optional user-supplied title."
required = false
+++
