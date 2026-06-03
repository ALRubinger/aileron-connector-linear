+++
name = "integration-gitlab-test-connection"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-gitlab-test-connection@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_gitlab_test_connection"]

[match]
intent = "Tests connectivity to a self-hosted GitLab instance and clears auth errors if successful."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_gitlab_test_connection"
idempotent = false

[approval]
required = true

[[inputs]]
name = "integrationId"
type = "String"
description = "The ID of the GitLab integration to test."
required = true
+++
