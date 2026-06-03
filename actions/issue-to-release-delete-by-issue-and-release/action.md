+++
name = "issue-to-release-delete-by-issue-and-release"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-to-release-delete-by-issue-and-release@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_to_release_delete_by_issue_and_release"]

[match]
intent = "Deletes an issue-to-release association by looking up the issue and release identifiers, removing the issue from the release."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_to_release_delete_by_issue_and_release"
idempotent = false

[approval]
required = true

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
