+++
name = "git-automation-target-branch-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/git-automation-target-branch-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["git_automation_target_branch_create"]

[match]
intent = "Creates a new Git target branch definition that scopes automation rules to pull requests targeting a specific branch pattern."

[[execute]]
id = "git"
connector = "github://ALRubinger/aileron-connector-linear"
op = "git_automation_target_branch_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "branchPattern"
type = "String"
description = "The target branch pattern."
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "isRegex"
type = "Boolean"
description = "Whether the branch pattern is a regular expression."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The team associated with the Git target branch automation."
required = true
+++
