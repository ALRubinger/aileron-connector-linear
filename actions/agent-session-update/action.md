+++
name = "agent-session-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-session-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_session_update"]

[match]
intent = "Updates an agent session."

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_session_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "addedExternalUrls"
type = "AgentSessionExternalUrlInput"
description = "URLs of external resources to be added to this session. Only updatable by the OAuth application that owns the session."
required = false

[[inputs]]
name = "dismissedAt"
type = "DateTime"
description = "[Internal] The time at which the agent session was dismissed. Set to null to un-dismiss. Only updatable by internal clients."
required = false

[[inputs]]
name = "externalLink"
type = "String"
description = "The URL of an external agent-hosted page associated with this session. Only updatable by the OAuth application that owns the session."
required = false

[[inputs]]
name = "externalUrls"
type = "AgentSessionExternalUrlInput"
description = "URLs of external resources associated with this session. Replaces existing URLs. Only updatable by the OAuth application that owns the session. If supplied, addedExternalUrls and removedExternalUrls are ignored."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the agent session to update."
required = true

[[inputs]]
name = "plan"
type = "JSONObject"
description = "A dynamically updated list of the agent's execution strategy. Only updatable by the OAuth application that owns the session."
required = false

[[inputs]]
name = "removedExternalUrls"
type = "String"
description = "URLs to be removed from this session. Only updatable by the OAuth application that owns the session."
required = false

[[inputs]]
name = "userState"
type = "AgentSessionUserStateInput"
description = "[Internal] User-specific state for the agent session. Only updatable by internal clients."
required = false
+++
