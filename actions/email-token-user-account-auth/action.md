+++
name = "email-token-user-account-auth"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/email-token-user-account-auth@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["email_token_user_account_auth"]

[match]
intent = "Authenticates a user account via email and authentication token."

[[execute]]
id = "email"
connector = "github://ALRubinger/aileron-connector-linear"
op = "email_token_user_account_auth"
idempotent = false

[approval]
required = true

[[inputs]]
name = "clientAuthCode"
type = "String"
description = "Auth code for the client initiating the login sequence."
required = false

[[inputs]]
name = "email"
type = "String"
description = "The email which to login via the magic login code."
required = true

[[inputs]]
name = "inviteLink"
type = "String"
description = "An optional invite link for a workspace."
required = false

[[inputs]]
name = "timezone"
type = "String"
description = "The timezone of the user's browser."
required = true

[[inputs]]
name = "token"
type = "String"
description = "The magic login code."
required = true
+++
