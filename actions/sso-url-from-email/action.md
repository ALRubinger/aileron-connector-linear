+++
name = "sso-url-from-email"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/sso-url-from-email@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["sso_url_from_email"]

[match]
intent = "Fetch SSO login URL for the email provided."

[[execute]]
id = "sso"
connector = "github://ALRubinger/aileron-connector-linear"
op = "sso_url_from_email"
idempotent = true

[[inputs]]
name = "email"
type = "String"
description = "Email to query the SSO login URL by."
required = true

[[inputs]]
name = "isDesktop"
type = "Boolean"
description = "Whether the client is the desktop app."
required = false

[[inputs]]
name = "type"
type = "IdentityProviderType"
description = "Type of identity provider."
required = true
+++
