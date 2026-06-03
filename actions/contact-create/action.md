+++
name = "contact-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/contact-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["contact_create"]

[match]
intent = "Creates a support contact message from an authenticated user. The message is saved and forwarded to Intercom for support handling."

[[execute]]
id = "contact"
connector = "github://ALRubinger/aileron-connector-linear"
op = "contact_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "browser"
type = "String"
description = "The user's browser name and version (e.g., 'Chrome 120')."
required = false

[[inputs]]
name = "clientVersion"
type = "String"
description = "The version of the Linear client application the user is running."
required = false

[[inputs]]
name = "device"
type = "String"
description = "The user's device type or model information."
required = false

[[inputs]]
name = "disappointmentRating"
type = "Int"
description = "How disappointed the user would be if they could no longer use Linear. Scale: 0 = not disappointed, 1 = somewhat disappointed, 2 = very disappointed, 3 = extremely disappointed."
required = false

[[inputs]]
name = "message"
type = "String"
description = "The feedback or support message submitted by the user."
required = true

[[inputs]]
name = "operatingSystem"
type = "String"
description = "The user's operating system name and version (e.g., 'macOS 14.0')."
required = false

[[inputs]]
name = "type"
type = "String"
description = "The type of support contact (e.g., bug report, feature request, general feedback)."
required = true
+++
