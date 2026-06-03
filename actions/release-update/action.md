+++
name = "release-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_update"]

[match]
intent = "Updates an existing release by ID. Supports updating name, description, version, commit SHA, pipeline, stage, and dates."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "commitSha"
type = "String"
description = "The commit SHA associated with this release."
required = false

[[inputs]]
name = "completedAt"
type = "DateTime"
description = "The time at which the release was completed."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the release."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the release to update."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The name of the release."
required = false

[[inputs]]
name = "stageId"
type = "String"
description = "The current stage of the release."
required = false

[[inputs]]
name = "startDate"
type = "TimelessDate"
description = "The estimated start date of the release."
required = false

[[inputs]]
name = "startedAt"
type = "DateTime"
description = "The time at which the release was started."
required = false

[[inputs]]
name = "targetDate"
type = "TimelessDate"
description = "The estimated completion date of the release."
required = false

[[inputs]]
name = "trashed"
type = "Boolean"
description = "Whether the release has been trashed."
required = false

[[inputs]]
name = "version"
type = "String"
description = "The version of the release."
required = false
+++
