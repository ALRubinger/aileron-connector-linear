+++
name = "agent-session-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-session-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_session_create"]

[match]
intent = "[Internal] Creates a new agent session on behalf of the current user"

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_session_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "appUserId"
type = "String"
description = "The app user (agent) to create a session for."
required = true

[[inputs]]
name = "context"
type = "JSONObject"
description = "[Internal] Serialized JSON representing the page contexts this session is related to. Used for direct chat sessions to provide context about the current page (e.g., Issue, Project)."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The issue that this session will be associated with. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = false

[[inputs]]
name = "pullRequestId"
type = "String"
description = "[Internal] Optional pull request to associate with the created session."
required = false
+++
