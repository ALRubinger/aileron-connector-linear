+++
name = "create-csv-export-report"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/create-csv-export-report@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["create_csv_export_report"]

[match]
intent = "Creates a CSV export report for the workspace. The report is generated asynchronously and delivered via email. Requires workspace admin export permission."

[[execute]]
id = "create"
connector = "github://ALRubinger/aileron-connector-linear"
op = "create_csv_export_report"
idempotent = false

[approval]
required = true

[[inputs]]
name = "includePrivateTeamIds"
type = "String"
description = "IDs of private teams to include in the export. If null, only data from public teams is exported."
required = false

[[inputs]]
name = "includeProtectedTeamIds"
type = "String"
description = "IDs of protected teams to include in the export. If null, protected team data is omitted."
required = false
+++
