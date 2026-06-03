+++
name = "document-content-history-timeline"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/document-content-history-timeline@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["document_content_history_timeline"]

[match]
intent = "[Internal] Fetches the fresh document content history timeline for the client history dialog."

[[execute]]
id = "document"
connector = "github://ALRubinger/aileron-connector-linear"
op = "document_content_history_timeline"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the document content to retrieve timeline entries for."
required = true
+++
