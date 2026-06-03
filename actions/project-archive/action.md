+++
name = "project-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_archive"]

[match]
intent = "Archives a project."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project to archive. Also the identifier from the URL is accepted."
required = true

[[inputs]]
name = "trash"
type = "Boolean"
description = "Whether to trash the project."
required = false
+++
