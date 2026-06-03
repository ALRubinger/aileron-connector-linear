+++
name = "project-unarchive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-unarchive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_unarchive"]

[match]
intent = "Restores a previously trashed or archived project."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_unarchive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project to restore. Also the identifier from the URL is accepted."
required = true
+++
