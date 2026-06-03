+++
name = "user-settings-flags-reset"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-settings-flags-reset@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_settings_flags_reset"]

[match]
intent = "Resets one or more of the user's setting flags. If no specific flags are provided, all flags are reset."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_settings_flags_reset"
idempotent = false

[approval]
required = true

[[inputs]]
name = "flags"
type = "UserFlagType"
description = "The flags to reset. If not provided all flags will be reset."
required = false
+++
