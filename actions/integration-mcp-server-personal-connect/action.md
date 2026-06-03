+++
name = "integration-mcp-server-personal-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-mcp-server-personal-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_mcp_server_personal_connect"]

[match]
intent = "[INTERNAL] Connects the user's personal account with an MCP server."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_mcp_server_personal_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "mcpServerDefinitionId"
type = "String"
description = "The workspace-approved MCP server definition used to start this connection."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The HTTP header name."
required = true

[[inputs]]
name = "serverUrl"
type = "String"
description = "The URL of the MCP server to connect with."
required = true

[[inputs]]
name = "value"
type = "String"
description = "The HTTP header value."
required = true
+++
