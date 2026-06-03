+++
name = "document-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/document-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["document_create"]

[match]
intent = "Creates a new document."

[[execute]]
id = "document"
connector = "github://ALRubinger/aileron-connector-linear"
op = "document_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The color of the icon."
required = false

[[inputs]]
name = "content"
type = "String"
description = "The document content as markdown."
required = false

[[inputs]]
name = "cycleId"
type = "String"
description = "[Internal] Related cycle for the document."
required = false

[[inputs]]
name = "icon"
type = "String"
description = "The icon of the document."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "[Internal] Related initiative for the document."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "Related issue for the document. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = false

[[inputs]]
name = "lastAppliedTemplateId"
type = "String"
description = "The ID of the last template applied to the document."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "Related project for the document."
required = false

[[inputs]]
name = "releaseId"
type = "String"
description = "Related release for the document."
required = false

[[inputs]]
name = "resourceFolderId"
type = "String"
description = "[Internal] The resource folder containing the document."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The order of the item in the resources list."
required = false

[[inputs]]
name = "subscriberIds"
type = "String"
description = "[INTERNAL] The identifiers of the users subscribing to this document."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "[Internal] Related team for the document."
required = false

[[inputs]]
name = "title"
type = "String"
description = "The title of the document."
required = true
+++
