+++
name = "user-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_update"]

[match]
intent = "Updates a user's profile information. Users can update their own profile; workspace admins can update any user's profile. SCIM-managed users may have restricted name changes."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "avatarUrl"
type = "String"
description = "The avatar image URL of the user."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The user description or a short bio."
required = false

[[inputs]]
name = "displayName"
type = "String"
description = "The display name of the user."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the user to update. Use `me` to reference currently authenticated user."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The name of the user."
required = false

[[inputs]]
name = "statusEmoji"
type = "String"
description = "The emoji part of the user status."
required = false

[[inputs]]
name = "statusLabel"
type = "String"
description = "The label part of the user status."
required = false

[[inputs]]
name = "statusUntilAt"
type = "DateTime"
description = "When the user status should be cleared."
required = false

[[inputs]]
name = "timezone"
type = "String"
description = "The local timezone of the user."
required = false

[[inputs]]
name = "title"
type = "String"
description = "The user's job title."
required = false
+++
