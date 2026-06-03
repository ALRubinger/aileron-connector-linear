+++
name = "create-initiative-update-reminder"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/create-initiative-update-reminder@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["create_initiative_update_reminder"]

[match]
intent = "Create a notification to remind a user about an initiative update."

[[execute]]
id = "create"
connector = "github://ALRubinger/aileron-connector-linear"
op = "create_initiative_update_reminder"
idempotent = false

[approval]
required = true

[[inputs]]
name = "initiativeId"
type = "String"
description = "The identifier of the initiative to remind about."
required = true

[[inputs]]
name = "userId"
type = "String"
description = "The user identifier to whom the notification will be sent. By default, it is set to the initiative owner."
required = false
+++
