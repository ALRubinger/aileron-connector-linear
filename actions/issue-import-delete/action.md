+++
name = "issue-import-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_delete"]

[match]
intent = "Deletes an import job."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "issueImportId"
type = "String"
description = "ID of the issue import to delete."
required = true
+++
