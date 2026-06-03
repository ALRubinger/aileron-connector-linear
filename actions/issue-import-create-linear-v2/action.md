+++
name = "issue-import-create-linear-v2"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-create-linear-v2@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_create_linear_v2"]

[match]
intent = "[INTERNAL] Kicks off a Linear to Linear import job."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_create_linear_v2"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "ID of issue import. If not provided it will be generated."
required = false

[[inputs]]
name = "linearSourceOrganizationId"
type = "String"
description = "The source workspace to import from."
required = true
+++
