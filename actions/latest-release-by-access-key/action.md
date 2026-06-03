+++
name = "latest-release-by-access-key"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/latest-release-by-access-key@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["latest_release_by_access_key"]

[match]
intent = "Returns the latest release for the pipeline associated with the access key."

[[execute]]
id = "latest"
connector = "github://ALRubinger/aileron-connector-linear"
op = "latest_release_by_access_key"
idempotent = true
+++
