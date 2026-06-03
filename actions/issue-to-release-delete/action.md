+++
name = "issue-to-release-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-to-release-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_to_release_delete"]

[match]
intent = "Deletes an issue-to-release association by its identifier, removing the issue from the release."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_to_release_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue-to-release association to delete."
required = true
+++
