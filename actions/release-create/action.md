+++
name = "release-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_create"]

[match]
intent = "Creates a new release in a pipeline. If no stage is specified, defaults to the first completed stage for continuous pipelines or the first started stage for scheduled pipelines."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_create"
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
description = "The time at which the release was completed (e.g. if importing from another system). Must be a time in the past and after createdAt. Cannot be provided with an incompatible stage."
required = false

[[inputs]]
name = "createdAt"
type = "DateTime"
description = "The time at which the release was created (e.g. if importing from another system). Must be a time in the past. If none is provided, the backend will generate the time as now."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the release."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the release."
required = true

[[inputs]]
name = "pipelineId"
type = "String"
description = "The identifier of the pipeline this release belongs to."
required = true

[[inputs]]
name = "stageId"
type = "String"
description = "The current stage of the release. Defaults to the first 'completed' stage for continuous pipelines, or the first 'started' stage for scheduled pipelines."
required = false

[[inputs]]
name = "startDate"
type = "TimelessDate"
description = "The estimated start date of the release."
required = false

[[inputs]]
name = "startedAt"
type = "DateTime"
description = "The time at which the release was started (e.g. if importing from another system). Must be a time in the past and after createdAt."
required = false

[[inputs]]
name = "targetDate"
type = "TimelessDate"
description = "The estimated completion date of the release."
required = false

[[inputs]]
name = "version"
type = "String"
description = "The version of the release."
required = false
+++
