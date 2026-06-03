+++
name = "release-pipeline-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-pipeline-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_pipeline_archive"]

[match]
intent = "Archives a release pipeline."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_pipeline_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the release pipeline to archive."
required = true
+++
