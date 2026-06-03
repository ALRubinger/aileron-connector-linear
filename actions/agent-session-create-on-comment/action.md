+++
name = "agent-session-create-on-comment"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-session-create-on-comment@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_session_create_on_comment"]

[match]
intent = "Creates a new agent session on a root comment."

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_session_create_on_comment"
idempotent = false

[approval]
required = true

[[inputs]]
name = "commentId"
type = "String"
description = "The root comment that this session will be associated with."
required = true

[[inputs]]
name = "externalLink"
type = "String"
description = "The URL of an external agent-hosted page associated with this session."
required = false

[[inputs]]
name = "externalUrls"
type = "AgentSessionExternalUrlInput"
description = "URLs of external resources associated with this session."
required = false
+++
