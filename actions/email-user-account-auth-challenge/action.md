+++
name = "email-user-account-auth-challenge"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/email-user-account-auth-challenge@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["email_user_account_auth_challenge"]

[match]
intent = "Finds or creates a new user account by email and sends an email with token."

[[execute]]
id = "email"
connector = "github://ALRubinger/aileron-connector-linear"
op = "email_user_account_auth_challenge"
idempotent = false

[approval]
required = true

[[inputs]]
name = "challengeResponse"
type = "String"
description = "Response from the login challenge."
required = false

[[inputs]]
name = "clientAuthCode"
type = "String"
description = "Auth code for the client initiating the sequence."
required = false

[[inputs]]
name = "email"
type = "String"
description = "The email for which to generate the magic login code."
required = true

[[inputs]]
name = "inviteLink"
type = "String"
description = "The workspace invite link to associate with this authentication."
required = false

[[inputs]]
name = "isDesktop"
type = "Boolean"
description = "Whether the login was requested from the desktop app."
required = false

[[inputs]]
name = "loginCodeOnly"
type = "Boolean"
description = "Whether to only return the login code. This is used by mobile apps to skip showing the login link."
required = false

[[inputs]]
name = "sessionId"
type = "String"
description = "PostHog session ID for attribution tracking."
required = false
+++
