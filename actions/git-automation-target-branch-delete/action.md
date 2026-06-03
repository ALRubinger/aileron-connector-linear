+++
name = "git-automation-target-branch-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/git-automation-target-branch-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["git_automation_target_branch_delete"]

[match]
intent = "Deletes a Git target branch definition and its associated automation rules."

[[execute]]
id = "git"
connector = "github://ALRubinger/aileron-connector-linear"
op = "git_automation_target_branch_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the Git target branch automation to archive."
required = true
+++
