+++
name = "attachment"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment"]

[match]
intent = "One specific issue attachment.\n[Deprecated] 'url' can no longer be used as the 'id' parameter. Use 'attachmentsForUrl' instead"

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
