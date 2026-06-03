+++
name = "release-update-by-pipeline-by-access-key"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-update-by-pipeline-by-access-key@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_update_by_pipeline_by_access_key"]

[match]
intent = "Updates a release by pipeline using an access key."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_update_by_pipeline_by_access_key"
idempotent = false

[approval]
required = true

[[inputs]]
name = "documents"
type = "ReleaseDocumentInput"
description = "Documents to attach to the updated release. Existing documents with the same title are updated."
required = false

[[inputs]]
name = "links"
type = "ReleaseLinkInput"
description = "External links to attach to the updated release."
required = false

[[inputs]]
name = "name"
type = "String"
description = "Optional release name to apply when updating the release."
required = false

[[inputs]]
name = "releaseNotes"
type = "ReleaseNoteInput"
description = "Release notes covering the updated release."
required = false

[[inputs]]
name = "stage"
type = "String"
description = "The stage name to set. First tries exact match, then falls back to case-insensitive matching with dashes/underscores treated as spaces."
required = false

[[inputs]]
name = "version"
type = "String"
description = "The version of the release to update. If not provided, the latest started or latest planned release will be updated."
required = false
+++
