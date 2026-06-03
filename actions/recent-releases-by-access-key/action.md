+++
name = "recent-releases-by-access-key"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/recent-releases-by-access-key@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["recent_releases_by_access_key"]

[match]
intent = "Returns recent in-progress and completed releases for the pipeline associated with the access key, ordered with in-progress first then most recently completed. Used by `linear-release` to walk candidates and pick the one whose `commitSha` is an ancestor of the current build commit, which disambiguates concurrent release trains on the same pipeline."

[[execute]]
id = "recent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "recent_releases_by_access_key"
idempotent = true

[[inputs]]
name = "limit"
type = "Int"
description = "Maximum number of releases to return. Defaults to 20 and is capped at 100."
required = false
+++
