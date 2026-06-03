+++
name = "email-intake-address-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/email-intake-address-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["email_intake_address_create"]

[match]
intent = "Creates a new email intake address."

[[execute]]
id = "email"
connector = "github://ALRubinger/aileron-connector-linear"
op = "email_intake_address_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "customerRequestsEnabled"
type = "Boolean"
description = "Whether customer requests are enabled."
required = false

[[inputs]]
name = "forwardingEmailAddress"
type = "String"
description = "The email address used to forward emails to the intake address."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "issueCanceledAutoReply"
type = "String"
description = "The auto-reply message for issue canceled."
required = false

[[inputs]]
name = "issueCanceledAutoReplyEnabled"
type = "Boolean"
description = "Whether the issue canceled auto-reply is enabled."
required = false

[[inputs]]
name = "issueCompletedAutoReply"
type = "String"
description = "The auto-reply message for issue completed."
required = false

[[inputs]]
name = "issueCompletedAutoReplyEnabled"
type = "Boolean"
description = "Whether the issue completed auto-reply is enabled."
required = false

[[inputs]]
name = "issueCreatedAutoReply"
type = "String"
description = "The auto-reply message for issue created."
required = false

[[inputs]]
name = "issueCreatedAutoReplyEnabled"
type = "Boolean"
description = "Whether the issue created auto-reply is enabled."
required = false

[[inputs]]
name = "reopenOnReply"
type = "Boolean"
description = "Whether to reopen completed or canceled issues when a substantive email reply is received."
required = false

[[inputs]]
name = "repliesEnabled"
type = "Boolean"
description = "Whether email replies are enabled."
required = false

[[inputs]]
name = "senderName"
type = "String"
description = "The name to be used for outgoing emails."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier or key of the team this email address will intake issues for."
required = false

[[inputs]]
name = "templateId"
type = "String"
description = "The identifier of the template this email address will intake issues for."
required = false

[[inputs]]
name = "type"
type = "EmailIntakeAddressType"
description = "The type of the email address. If not provided, the backend will default to team or template."
required = false

[[inputs]]
name = "useUserNamesInReplies"
type = "Boolean"
description = "Whether the commenter's name is included in the email replies."
required = false
+++
