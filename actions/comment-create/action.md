+++
name = "comment-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/comment-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["create_comment"]

[match]
intent = "Creates a new comment."

[[execute]]
id = "comment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "comment_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "body"
type = "String"
description = "The comment content in markdown format."
required = false

[[inputs]]
name = "bodyData"
type = "JSON"
description = "[Internal] The comment content as a Prosemirror document."
required = false

[[inputs]]
name = "createAsUser"
type = "String"
description = "Create comment as a user with the provided name. This option is only available to OAuth applications creating comments in `actor=app` mode."
required = false

[[inputs]]
name = "createOnSyncedSlackThread"
type = "Boolean"
description = "Flag to indicate this comment should be created on the issue's synced Slack comment thread. If no synced Slack comment thread exists, the mutation will fail. If there are multiple synced Slack threads on the issue, the oldest one will be targeted."
required = false

[[inputs]]
name = "createdAt"
type = "DateTime"
description = "The time at which the comment was created (e.g. if importing from another system). Must be a time in the past. If none is provided, the backend will generate the time as now."
required = false

[[inputs]]
name = "displayIconUrl"
type = "String"
description = "Provide an external user avatar URL. Can only be used in conjunction with the `createAsUser` options. This option is only available to OAuth applications creating comments in `actor=app` mode."
required = false

[[inputs]]
name = "doNotSubscribeToIssue"
type = "Boolean"
description = "Flag to prevent auto subscription to the issue the comment is created on."
required = false

[[inputs]]
name = "documentContentId"
type = "String"
description = "The document content to associate the comment with."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "The initiative to associate the comment with."
required = false

[[inputs]]
name = "initiativeUpdateId"
type = "String"
description = "The initiative update to associate the comment with."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The issue to associate the comment with. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = false

[[inputs]]
name = "parentId"
type = "String"
description = "The parent comment under which to nest a current comment."
required = false

[[inputs]]
name = "postId"
type = "String"
description = "The post to associate the comment with."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "The project to associate the comment with."
required = false

[[inputs]]
name = "projectUpdateId"
type = "String"
description = "The project update to associate the comment with."
required = false

[[inputs]]
name = "quotedText"
type = "String"
description = "The text that this comment references. Only defined for inline comments."
required = false

[[inputs]]
name = "subscriberIds"
type = "String"
description = "[INTERNAL] The identifiers of the users subscribing to this comment thread."
required = false
+++
