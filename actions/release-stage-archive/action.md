+++
name = "release-stage-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-stage-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_stage_archive"]

[match]
intent = "Archives a release stage. Only started-type stages can be archived, and only if they have no active releases and at least one other stage of the same type remains. Cannot archive the last non-frozen started stage."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_stage_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the release stage to archive."
required = true
+++
