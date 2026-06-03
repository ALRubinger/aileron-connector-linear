+++
name = "issue-import-process"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-process@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_process"]

[match]
intent = "Kicks off import processing."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_process"
idempotent = false

[approval]
required = true

[[inputs]]
name = "issueImportId"
type = "String"
description = "ID of the issue import which we're going to process."
required = true

[[inputs]]
name = "mapping"
type = "JSONObject"
description = "The mapping configuration to use for processing the import."
required = true
+++
