+++
name = "integration-gitlab-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-gitlab-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_gitlab_connect"]

[match]
intent = "Connects the workspace with a GitLab Access Token."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_gitlab_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "accessToken"
type = "String"
description = "The GitLab Access Token to connect with."
required = true

[[inputs]]
name = "expiresAt"
type = "String"
description = "ISO timestamp when the access token expires. Used as a fallback when `validationProjectPath` is set and `personal_access_tokens/self` is unreachable."
required = false

[[inputs]]
name = "gitlabUrl"
type = "String"
description = "The URL of the GitLab installation."
required = true

[[inputs]]
name = "readonly"
type = "Boolean"
description = "Whether the access token is limited to a read-only scope. Used as a fallback when `validationProjectPath` is set and `personal_access_tokens/self` is unreachable."
required = false

[[inputs]]
name = "validationProjectPath"
type = "String"
description = "Optional path or numeric ID of a project to use for the setup health check. Set this when the GitLab tenant blocks non-project API endpoints."
required = false
+++
