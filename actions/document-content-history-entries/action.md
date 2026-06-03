+++
name = "document-content-history-entries"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/document-content-history-entries@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["document_content_history_entries"]

[match]
intent = "[Internal] Fetches document content history entries by their IDs, including content data."

[[execute]]
id = "document"
connector = "github://ALRubinger/aileron-connector-linear"
op = "document_content_history_entries"
idempotent = true

[[inputs]]
name = "entryIds"
type = "String"
description = "The identifiers of the history entries to retrieve. Maximum 10 per request."
required = true
+++
