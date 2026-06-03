+++
name = "release-pipeline-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-pipeline-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_pipeline_update"]

[match]
intent = "Updates an existing release pipeline. Supports updating name, slug, type, production flag, path patterns, and team associations. Private teams that the current user cannot access are preserved in the team list."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_pipeline_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "autoGenerateReleaseNotesOnCompletion"
type = "Boolean"
description = "Whether to automatically generate a release note when a release is completed. Only applies to scheduled pipelines; ignored for continuous pipelines."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the release pipeline to update."
required = true

[[inputs]]
name = "includePathPatterns"
type = "String"
description = "Glob patterns to include commits affecting matching file paths."
required = false

[[inputs]]
name = "isProduction"
type = "Boolean"
description = "Whether this pipeline targets a production environment. Default to true."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the pipeline."
required = false

[[inputs]]
name = "slugId"
type = "String"
description = "The pipeline's unique slug identifier."
required = false

[[inputs]]
name = "teamIds"
type = "String"
description = "The identifiers of the teams this pipeline is associated with."
required = false

[[inputs]]
name = "type"
type = "ReleasePipelineType"
description = "The type of the pipeline."
required = false
+++
