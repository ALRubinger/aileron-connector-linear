+++
name = "issue-relation-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-relation-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_relation_create"]

[match]
intent = "Creates a new issue relation."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_relation_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The identifier of the issue that is related to another issue. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = true

[[inputs]]
name = "overrideCreatedAt"
type = "DateTime"
description = "Used by client undo operations. Should not be set directly."
required = false

[[inputs]]
name = "relatedIssueId"
type = "String"
description = "The identifier of the related issue. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = true

[[inputs]]
name = "type"
type = "IssueRelationType"
description = "The type of relation of the issue to the related issue."
required = true
+++
