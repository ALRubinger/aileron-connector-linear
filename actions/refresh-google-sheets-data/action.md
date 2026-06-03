+++
name = "refresh-google-sheets-data"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/refresh-google-sheets-data@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["refresh_google_sheets_data"]

[match]
intent = "Manually update Google Sheets data."

[[execute]]
id = "refresh"
connector = "github://ALRubinger/aileron-connector-linear"
op = "refresh_google_sheets_data"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the Google Sheets integration to update."
required = true

[[inputs]]
name = "type"
type = "String"
description = "The type of export."
required = false
+++
