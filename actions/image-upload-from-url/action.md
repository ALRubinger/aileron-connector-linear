+++
name = "image-upload-from-url"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/image-upload-from-url@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["image_upload_from_url"]

[match]
intent = "Upload an image from an URL to Linear."

[[execute]]
id = "image"
connector = "github://ALRubinger/aileron-connector-linear"
op = "image_upload_from_url"
idempotent = false

[approval]
required = true

[[inputs]]
name = "url"
type = "String"
description = "URL of the file to be uploaded to Linear."
required = true
+++
