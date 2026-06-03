+++
name = "organization-exists"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-exists@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_exists"]

[match]
intent = "Checks whether a workspace with the given URL key exists."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_exists"
idempotent = true

[[inputs]]
name = "urlKey"
type = "String"
description = "The URL key of the workspace to check."
required = true
+++
