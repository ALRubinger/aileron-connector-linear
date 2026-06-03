+++
name = "issue-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_updated"]

[match]
intent = "Triggered when an issue is updated"

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_updated"
idempotent = true

[[inputs]]
name = "assigneeId"
type = "IDComparator"
description = "Filter by assignee ID."
required = false

[[inputs]]
name = "parentId"
type = "IDComparator"
description = "Filter by parent issue ID."
required = false

[[inputs]]
name = "projectId"
type = "IDComparator"
description = "Filter by project ID."
required = false

[[inputs]]
name = "stateId"
type = "IDComparator"
description = "Filter by workflow state ID."
required = false

[[inputs]]
name = "teamId"
type = "IDComparator"
description = "Filter by team ID."
required = false
+++
