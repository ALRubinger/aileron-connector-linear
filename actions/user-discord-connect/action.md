+++
name = "user-discord-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-discord-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_discord_connect"]

[match]
intent = "Connects the Discord user to this Linear account via OAuth2."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_discord_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Discord OAuth code."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Discord OAuth redirect URI."
required = true
+++
