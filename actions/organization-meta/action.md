+++
name = "organization-meta"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/organization-meta@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["organization_meta"]

[match]
intent = "[INTERNAL] Get workspace metadata by URL key or workspace ID."

[[execute]]
id = "organization"
connector = "github://ALRubinger/aileron-connector-linear"
op = "organization_meta"
idempotent = true

[[inputs]]
name = "urlKey"
type = "String"
description = "The URL key of the workspace to retrieve metadata for."
required = true
+++
