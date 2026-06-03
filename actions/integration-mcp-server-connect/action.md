+++
name = "integration-mcp-server-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-mcp-server-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_mcp_server_connect"]

[match]
intent = "[INTERNAL] Connects the workspace with an MCP server."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_mcp_server_connect"
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
name = "teamId"
type = "String"
description = "The ID of the team to connect the MCP server to."
required = false

[[inputs]]
name = "value"
type = "String"
description = "The HTTP header value."
required = true

[[inputs]]
name = "workflowDefinitionDraftId"
type = "String"
description = "The ID of the workflow definition draft to connect the MCP server to."
required = false

[[inputs]]
name = "workflowDefinitionId"
type = "String"
description = "The ID of the workflow definition to connect the MCP server to."
required = false
+++
