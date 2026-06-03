+++
name = "ai-conversation-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/ai-conversation-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["ai_conversation_updated"]

[match]
intent = "Triggered when an ai conversation is updated"

[[execute]]
id = "ai"
connector = "github://ALRubinger/aileron-connector-linear"
op = "ai_conversation_updated"
idempotent = true
+++
