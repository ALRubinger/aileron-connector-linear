+++
name = "release-pipeline-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-pipeline-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_pipeline_delete"]

[match]
intent = "Moves a release pipeline to the trash bin. Trashed pipelines are archived and will be permanently deleted after a retention period, along with all their releases. If the pipeline is already archived, it is marked as trashed with a fresh archive timestamp."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_pipeline_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the release pipeline to delete."
required = true
+++
