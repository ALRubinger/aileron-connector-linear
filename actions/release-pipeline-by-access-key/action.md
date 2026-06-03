+++
name = "release-pipeline-by-access-key"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-pipeline-by-access-key@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_pipeline_by_access_key"]

[match]
intent = "Returns a release pipeline by ID. Requires the access key to have access to the pipeline."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_pipeline_by_access_key"
idempotent = true
+++
