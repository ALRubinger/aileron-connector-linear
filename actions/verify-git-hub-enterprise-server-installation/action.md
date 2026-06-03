+++
name = "verify-git-hub-enterprise-server-installation"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/verify-git-hub-enterprise-server-installation@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["verify_git_hub_enterprise_server_installation"]

[match]
intent = "Verify that we received the correct response from the GitHub Enterprise Server."

[[execute]]
id = "verify"
connector = "github://ALRubinger/aileron-connector-linear"
op = "verify_git_hub_enterprise_server_installation"
idempotent = true

[[inputs]]
name = "integrationId"
type = "String"
description = "The integration ID."
required = true
+++
