+++
name = "issue-import-check-sync"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-check-sync@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_check_sync"]

[match]
intent = "Checks whether it will be possible to set up sync for this project or repository at the end of import."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_check_sync"
idempotent = true

[[inputs]]
name = "issueImportId"
type = "String"
description = "The ID of the issue import for which to check sync eligibility"
required = true
+++
