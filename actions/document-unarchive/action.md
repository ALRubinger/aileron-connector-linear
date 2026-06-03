+++
name = "document-unarchive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/document-unarchive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["document_unarchive"]

[match]
intent = "Restores a previously trashed document by unarchiving it."

[[execute]]
id = "document"
connector = "github://ALRubinger/aileron-connector-linear"
op = "document_unarchive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the document to restore."
required = true
+++
