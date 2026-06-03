+++
name = "document"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/document@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["document"]

[match]
intent = "A specific document by ID or slug."

[[execute]]
id = "document"
connector = "github://ALRubinger/aileron-connector-linear"
op = "document"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier or slug of the document to retrieve."
required = true
+++
