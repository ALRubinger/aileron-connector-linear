+++
name = "import-file-upload"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/import-file-upload@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["import_file_upload"]

[match]
intent = "XHR request payload to upload a file for import, directly to Linear's cloud storage."

[[execute]]
id = "import"
connector = "github://ALRubinger/aileron-connector-linear"
op = "import_file_upload"
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
