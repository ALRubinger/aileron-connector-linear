+++
name = "document-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/document-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["document_delete"]

[match]
intent = "Deletes (trashes) a document. The document is marked as trashed and archived, but not permanently removed."

[[execute]]
id = "document"
connector = "github://ALRubinger/aileron-connector-linear"
op = "document_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the document to delete."
required = true
+++
