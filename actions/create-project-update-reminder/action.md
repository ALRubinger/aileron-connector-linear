+++
name = "create-project-update-reminder"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/create-project-update-reminder@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["create_project_update_reminder"]

[match]
intent = "Create a notification to remind a user about a project update."

[[execute]]
id = "create"
connector = "github://ALRubinger/aileron-connector-linear"
op = "create_project_update_reminder"
idempotent = false

[approval]
required = true

[[inputs]]
name = "projectId"
type = "String"
description = "The identifier of the project to remind about."
required = true

[[inputs]]
name = "userId"
type = "String"
description = "The user identifier to whom the notification will be sent. By default, it is set to the project lead."
required = false
+++
