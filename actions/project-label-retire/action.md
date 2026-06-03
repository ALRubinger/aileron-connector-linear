+++
name = "project-label-retire"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-label-retire@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_label_retire"]

[match]
intent = "Retires a project label. Retired labels remain on existing projects but cannot be applied to new ones."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_label_retire"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the label to retire."
required = true
+++
