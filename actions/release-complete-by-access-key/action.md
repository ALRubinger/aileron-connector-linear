+++
name = "release-complete-by-access-key"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-complete-by-access-key@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_complete_by_access_key"]

[match]
intent = "Marks a release as completed using an access key. If version is provided, completes that specific release; otherwise completes the most recent started release. The pipeline is inferred from the access key."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_complete_by_access_key"
idempotent = false

[approval]
required = true

[[inputs]]
name = "commitSha"
type = "String"
description = "The commit SHA associated with this completion. If a completed release with this SHA already exists, it will be returned instead of completing a new release."
required = false

[[inputs]]
name = "documents"
type = "ReleaseDocumentInput"
description = "Documents to attach to the completed release. Existing documents with the same title are updated."
required = false

[[inputs]]
name = "links"
type = "ReleaseLinkInput"
description = "External links to attach to the completed release."
required = false

[[inputs]]
name = "name"
type = "String"
description = "Optional release name to apply when completing the release."
required = false

[[inputs]]
name = "releaseNotes"
type = "ReleaseNoteInput"
description = "Release notes covering the completed release."
required = false

[[inputs]]
name = "version"
type = "String"
description = "The version of the release to complete. If not provided, the latest started release will be completed."
required = false
+++
