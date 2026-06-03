+++
name = "project-external-sync-disable"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-external-sync-disable@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_external_sync_disable"]

[match]
intent = "Disables external sync on a project."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_external_sync_disable"
idempotent = false

[approval]
required = true

[[inputs]]
name = "projectId"
type = "String"
description = "The ID of the project to disable external sync for."
required = true

[[inputs]]
name = "syncSource"
type = "ExternalSyncService"
description = "The source of the external sync to disable."
required = true
+++
