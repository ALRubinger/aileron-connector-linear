+++
name = "document-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/document-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["document_created"]

[match]
intent = "Triggered when a document is created"

[[execute]]
id = "document"
connector = "github://ALRubinger/aileron-connector-linear"
op = "document_created"
idempotent = true
+++
