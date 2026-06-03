+++
name = "file-upload"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/file-upload@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["file_upload"]

[match]
intent = "XHR request payload to upload an images, video and other attachments directly to Linear's cloud storage."

[[execute]]
id = "file"
connector = "github://ALRubinger/aileron-connector-linear"
op = "file_upload"
idempotent = false

[approval]
required = true

[[inputs]]
name = "contentType"
type = "String"
description = "MIME type of the uploaded file."
required = true

[[inputs]]
name = "filename"
type = "String"
description = "Filename of the uploaded file."
required = true

[[inputs]]
name = "makePublic"
type = "Boolean"
description = "Should the file be made publicly accessible (default: false)."
required = false

[[inputs]]
name = "metaData"
type = "JSON"
description = "Optional metadata."
required = false

[[inputs]]
name = "size"
type = "Int"
description = "File size of the uploaded file."
required = true
+++
