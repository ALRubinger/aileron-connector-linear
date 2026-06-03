+++
name = "oauth-applications"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/oauth-applications@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["oauth_applications"]

[match]
intent = "[ALPHA] OAuth applications created by the calling OAuth application through the public API."

[[execute]]
id = "oauth"
connector = "github://ALRubinger/aileron-connector-linear"
op = "oauth_applications"
idempotent = true
+++
