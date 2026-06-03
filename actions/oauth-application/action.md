+++
name = "oauth-application"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/oauth-application@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["oauth_application"]

[match]
intent = "[ALPHA] A specific OAuth application created by the calling OAuth application through the public API."

[[execute]]
id = "oauth"
connector = "github://ALRubinger/aileron-connector-linear"
op = "oauth_application"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the OAuth application to retrieve."
required = true
+++
