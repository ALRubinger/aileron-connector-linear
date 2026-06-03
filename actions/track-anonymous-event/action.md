+++
name = "track-anonymous-event"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/track-anonymous-event@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["track_anonymous_event"]

[match]
intent = "Track an anonymous analytics event."

[[execute]]
id = "track"
connector = "github://ALRubinger/aileron-connector-linear"
op = "track_anonymous_event"
idempotent = false

[approval]
required = true

[[inputs]]
name = "event"
type = "String"
description = "The event name to track."
required = true

[[inputs]]
name = "properties"
type = "JSONObject"
description = "Optional properties for the event."
required = false

[[inputs]]
name = "sessionId"
type = "String"
description = "Client session ID for PostHog session correlation."
required = false
+++
