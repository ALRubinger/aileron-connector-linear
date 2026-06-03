+++
name = "notification"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification"]

[match]
intent = "A specific notification by ID."

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the notification to retrieve."
required = true
+++
