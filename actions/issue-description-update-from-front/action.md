+++
name = "issue-description-update-from-front"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-description-update-from-front@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_description_update_from_front"]

[match]
intent = "[INTERNAL] Updates an issue description from the Front app to handle Front attachments correctly."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_description_update_from_front"
idempotent = false

[approval]
required = true

[[inputs]]
name = "description"
type = "String"
description = "Description to update the issue with."
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to update."
required = true
+++
