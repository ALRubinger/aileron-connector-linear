+++
name = "attachment-issue"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-issue@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_issue"]

[match]
intent = "Query an issue by its associated attachment, and its id."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_issue"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "`id` of the attachment for which you'll want to get the issue for. [Deprecated] `url` as the `id` parameter."
required = true
+++
