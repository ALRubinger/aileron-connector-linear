+++
name = "initiative-add-label"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-add-label@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_add_label"]

[match]
intent = "[Internal]Adds a label to an initiative."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_add_label"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the initiative to add the label to."
required = true

[[inputs]]
name = "labelId"
type = "String"
description = "The identifier of the label to add."
required = true
+++
