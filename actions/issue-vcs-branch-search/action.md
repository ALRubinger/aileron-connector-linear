+++
name = "issue-vcs-branch-search"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-vcs-branch-search@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_vcs_branch_search"]

[match]
intent = "Find issue based on the VCS branch name."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_vcs_branch_search"
idempotent = true

[[inputs]]
name = "branchName"
type = "String"
description = "The VCS branch name to search for."
required = true
+++
