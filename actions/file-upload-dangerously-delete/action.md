+++
name = "file-upload-dangerously-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/file-upload-dangerously-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["file_upload_dangerously_delete"]

[match]
intent = "[INTERNAL] Permanently delete an uploaded file by asset URL. This should be used as a last resort and will break comments and documents that reference the asset."

[[execute]]
id = "file"
connector = "github://ALRubinger/aileron-connector-linear"
op = "file_upload_dangerously_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "assetUrl"
type = "String"
description = "The asset URL of the uploaded file to delete."
required = true
+++
