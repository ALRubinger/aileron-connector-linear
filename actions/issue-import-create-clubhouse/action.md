+++
name = "issue-import-create-clubhouse"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-create-clubhouse@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_create_clubhouse"]

[match]
intent = "Kicks off a Shortcut (formerly Clubhouse) import job."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_create_clubhouse"
idempotent = false

[approval]
required = true

[[inputs]]
name = "clubhouseGroupName"
type = "String"
description = "Shortcut (formerly Clubhouse) group name to choose which issues we should import."
required = true

[[inputs]]
name = "clubhouseToken"
type = "String"
description = "Shortcut (formerly Clubhouse) token to fetch information from the Clubhouse API."
required = true

[[inputs]]
name = "id"
type = "String"
description = "ID of issue import. If not provided it will be generated."
required = false

[[inputs]]
name = "includeClosedIssues"
type = "Boolean"
description = "Whether or not we should collect the data for closed issues."
required = false

[[inputs]]
name = "instantProcess"
type = "Boolean"
description = "Whether to instantly process the import with the default configuration mapping."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "ID of the team into which to import data."
required = false

[[inputs]]
name = "teamName"
type = "String"
description = "Name of new team. When teamId is not set."
required = false
+++
