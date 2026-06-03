+++
name = "reaction-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/reaction-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["reaction_create"]

[match]
intent = "Creates a new reaction."

[[execute]]
id = "reaction"
connector = "github://ALRubinger/aileron-connector-linear"
op = "reaction_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "commentId"
type = "String"
description = "The comment to associate the reaction with."
required = false

[[inputs]]
name = "emoji"
type = "String"
description = "The emoji the user reacted with."
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeUpdateId"
type = "String"
description = "The update to associate the reaction with."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The issue to associate the reaction with. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = false

[[inputs]]
name = "postId"
type = "String"
description = "[Internal] The post to associate the reaction with."
required = false

[[inputs]]
name = "projectUpdateId"
type = "String"
description = "The project update to associate the reaction with."
required = false

[[inputs]]
name = "pullRequestCommentId"
type = "String"
description = "[Internal] The pull request comment to associate the reaction with."
required = false

[[inputs]]
name = "pullRequestId"
type = "String"
description = "[Internal] The pull request to associate the reaction with."
required = false
+++
