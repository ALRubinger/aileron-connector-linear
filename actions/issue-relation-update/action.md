+++
name = "issue-relation-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-relation-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_relation_update"]

[match]
intent = "Updates an issue relation."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_relation_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue relation to update."
required = true

[[inputs]]
name = "issueId"
type = "String"
description = "The identifier of the issue that is related to another issue. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = false

[[inputs]]
name = "relatedIssueId"
type = "String"
description = "The identifier of the related issue. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = false

[[inputs]]
name = "type"
type = "String"
description = "The type of relation of the issue to the related issue."
required = false
+++
