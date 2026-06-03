+++
name = "release-unarchive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-unarchive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_unarchive"]

[match]
intent = "Unarchives a release."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_unarchive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the release to unarchive."
required = true
+++
