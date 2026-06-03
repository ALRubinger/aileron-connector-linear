+++
name = "issue-import-check-c-s-v"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-check-c-s-v@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_check_c_s_v"]

[match]
intent = "Checks a CSV file validity against a specific import service."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_check_c_s_v"
idempotent = true

[[inputs]]
name = "csvUrl"
type = "String"
description = "The CSV file storage URL."
required = true

[[inputs]]
name = "service"
type = "String"
description = "The service the CSV contains data from."
required = true
+++
