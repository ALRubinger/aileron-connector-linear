+++
name = "release-stage-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-stage-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_stage_create"]

[match]
intent = "Creates a new release stage in a pipeline. Non-started stages must use default names and colors, and only one stage of each non-started type is allowed per pipeline. Started stages can optionally be frozen, but at least one non-frozen started stage must remain."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_stage_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The UI color of the stage as a HEX string."
required = true

[[inputs]]
name = "frozen"
type = "Boolean"
description = "Whether this stage is frozen. Only applicable to started stages."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the stage."
required = true

[[inputs]]
name = "pipelineId"
type = "String"
description = "The identifier of the pipeline this stage belongs to."
required = true

[[inputs]]
name = "position"
type = "Float"
description = "The position of the stage."
required = true

[[inputs]]
name = "type"
type = "ReleaseStageType"
description = "The type of the stage."
required = true
+++
