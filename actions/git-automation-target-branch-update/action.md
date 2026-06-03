+++
name = "git-automation-target-branch-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/git-automation-target-branch-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["git_automation_target_branch_update"]

[match]
intent = "Updates an existing Git target branch definition, including its branch pattern and regex flag."

[[execute]]
id = "git"
connector = "github://ALRubinger/aileron-connector-linear"
op = "git_automation_target_branch_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "branchPattern"
type = "String"
description = "The target branch pattern."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the Git target branch automation to update."
required = true

[[inputs]]
name = "isRegex"
type = "Boolean"
description = "Whether the branch pattern is a regular expression."
required = false
+++
