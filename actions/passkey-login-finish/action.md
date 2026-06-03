+++
name = "passkey-login-finish"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/passkey-login-finish@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["passkey_login_finish"]

[match]
intent = "[INTERNAL] Finish passkey login process."

[[execute]]
id = "passkey"
connector = "github://ALRubinger/aileron-connector-linear"
op = "passkey_login_finish"
idempotent = false

[approval]
required = true

[[inputs]]
name = "authId"
type = "String"
description = "Random ID to start passkey login with."
required = true

[[inputs]]
name = "response"
type = "JSONObject"
description = "The WebAuthn authenticator response from the passkey challenge."
required = true
+++
