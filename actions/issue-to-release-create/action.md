+++
name = "issue-to-release-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-to-release-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_to_release_create"]

[match]
intent = "Creates a new association between an issue and a release, linking the issue to the release for tracking purposes."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_to_release_create"
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
description = "The identifier of the issue. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = true

[[inputs]]
name = "releaseId"
type = "String"
description = "The identifier of the release."
required = true
+++
