+++
name = "oauth-application-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/oauth-application-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["oauth_application_archive"]

[match]
intent = "[ALPHA] Archives an OAuth application created by the calling OAuth application."

[[execute]]
id = "oauth"
connector = "github://ALRubinger/aileron-connector-linear"
op = "oauth_application_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the OAuth application to archive."
required = true
+++
