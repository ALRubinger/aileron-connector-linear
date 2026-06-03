+++
name = "release-sync"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-sync@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_sync"]

[match]
intent = "Syncs release data by resolving issue and pull request references and associating them with a release. For continuous pipelines, creates a new completed release. For scheduled pipelines, finds or creates a started release and accumulates issues into it."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_sync"
idempotent = false

[approval]
required = true

[[inputs]]
name = "commitSha"
type = "String"
description = "The commit SHA associated with this release."
required = true

[[inputs]]
name = "debugSink"
type = "ReleaseDebugSinkInput"
description = "Debug information for release creation diagnostics."
required = false

[[inputs]]
name = "documents"
type = "ReleaseDocumentInput"
description = "Documents to attach to the release. Existing documents on the release with the same title are updated."
required = false

[[inputs]]
name = "issueReferences"
type = "IssueReferenceInput"
description = "Issue references (e.g. ENG-123) to associate with this release."
required = false

[[inputs]]
name = "links"
type = "ReleaseLinkInput"
description = "External links to attach to the release."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the release."
required = false

[[inputs]]
name = "pipelineId"
type = "String"
description = "The identifier of the pipeline this release belongs to."
required = true

[[inputs]]
name = "pullRequestReferences"
type = "PullRequestReferenceInput"
description = "Pull request references to look up. Issues linked to found PRs will be associated with this release."
required = false

[[inputs]]
name = "releaseNotes"
type = "ReleaseNoteInput"
description = "Release notes covering this release. Upserts the existing release notes that cover only this release."
required = false

[[inputs]]
name = "repository"
type = "RepositoryDataInput"
description = "Information about the source repository."
required = false

[[inputs]]
name = "revertedIssueReferences"
type = "IssueReferenceInput"
description = "Issue references that were reverted and should be removed from the release."
required = false

[[inputs]]
name = "version"
type = "String"
description = "The version of the release."
required = false
+++
