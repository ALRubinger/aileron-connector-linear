+++
name = "document-content-history"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/document-content-history@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["document_content_history"]

[match]
intent = "A collection of document content history entries."

[[execute]]
id = "document"
connector = "github://ALRubinger/aileron-connector-linear"
op = "document_content_history"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the document content to retrieve history for."
required = true
+++
