+++
name = "issue-import-create-github"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-create-github@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_create_github"]

[match]
intent = "Kicks off a GitHub import job."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_create_github"
idempotent = false

[approval]
required = true

[[inputs]]
name = "githubLabels"
type = "String"
description = "Labels to use to filter the import data. Only issues matching any of these filters will be imported."
required = false

[[inputs]]
name = "githubRepoIds"
type = "Int"
description = "IDs of the Github repositories from which we will import data."
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
