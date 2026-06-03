+++
name = "document-content-draft-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/document-content-draft-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["document_content_draft_updated"]

[match]
intent = "Triggered when a document content draft is updated"

[[execute]]
id = "document"
connector = "github://ALRubinger/aileron-connector-linear"
op = "document_content_draft_updated"
idempotent = true
+++
