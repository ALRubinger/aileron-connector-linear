+++
name = "google-user-account-auth"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/google-user-account-auth@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["google_user_account_auth"]

[match]
intent = "Authenticate user account through Google OAuth. This is the 2nd step of OAuth flow."

[[execute]]
id = "google"
connector = "github://ALRubinger/aileron-connector-linear"
op = "google_user_account_auth"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "Code returned from Google's OAuth flow."
required = true

[[inputs]]
name = "disallowSignup"
type = "Boolean"
description = "An optional parameter to disable new user signup and force login. Default: false."
required = false

[[inputs]]
name = "inviteLink"
type = "String"
description = "An optional invite link for a workspace used to populate available workspaces."
required = false

[[inputs]]
name = "redirectUri"
type = "String"
description = "The URI to redirect the user to."
required = false

[[inputs]]
name = "sessionId"
type = "String"
description = "PostHog session ID for attribution tracking."
required = false

[[inputs]]
name = "timezone"
type = "String"
description = "The timezone of the user's browser."
required = true
+++
