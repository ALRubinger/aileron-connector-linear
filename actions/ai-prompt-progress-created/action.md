+++
name = "ai-prompt-progress-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/ai-prompt-progress-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["ai_prompt_progress_created"]

[match]
intent = "Triggered when an ai prompt progress is created"

[[execute]]
id = "ai"
connector = "github://ALRubinger/aileron-connector-linear"
op = "ai_prompt_progress_created"
idempotent = true

[[inputs]]
name = "commentId"
type = "IDComparator"
description = "[Internal] Filter by comment ID."
required = false

[[inputs]]
name = "issueId"
type = "IDComparator"
description = "[Internal] Filter by issue ID."
required = false

[[inputs]]
name = "pullRequestCommentId"
type = "IDComparator"
description = "[Internal] Filter by pull request comment ID."
required = false

[[inputs]]
name = "status"
type = "AiPromptProgressStatusComparator"
description = "[Internal] Filter by prompt workflow status."
required = false

[[inputs]]
name = "type"
type = "AiPromptTypeComparator"
description = "[Internal] Filter by prompt workflow type."
required = false
+++
