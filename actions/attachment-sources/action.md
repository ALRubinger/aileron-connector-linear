+++
name = "attachment-sources"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-sources@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_sources"]

[match]
intent = "[Internal] Get a list of all unique attachment sources in the workspace."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_sources"
idempotent = true

[[inputs]]
name = "teamId"
type = "String"
description = "(optional) if provided will only return attachment sources for the given team."
required = false
+++
