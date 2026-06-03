+++
name = "agent-session-update-external-url"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-session-update-external-url@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_session_update_external_url"]

[match]
intent = "Updates the externalUrl of an agent session, which is an agent-hosted page associated with this session."

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_session_update_external_url"
idempotent = false

[approval]
required = true

[[inputs]]
name = "addedExternalUrls"
type = "AgentSessionExternalUrlInput"
description = "URLs of external resources to be added to this session."
required = false

[[inputs]]
name = "externalLink"
type = "String"
description = "The URL of an external agent-hosted page associated with this session."
required = false

[[inputs]]
name = "externalUrls"
type = "AgentSessionExternalUrlInput"
description = "URLs of external resources associated with this session. Replaces existing URLs."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the agent session to update."
required = true

[[inputs]]
name = "removedExternalUrls"
type = "String"
description = "URLs to be removed from this session."
required = false
+++
