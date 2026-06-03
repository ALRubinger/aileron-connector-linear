+++
name = "custom-view-details-suggestion"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/custom-view-details-suggestion@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["custom_view_details_suggestion"]

[match]
intent = "[INTERNAL] Suggests metadata (name, description, icon) for a view based on its filters using AI. Rate-limited to 30 requests per minute."

[[execute]]
id = "custom"
connector = "github://ALRubinger/aileron-connector-linear"
op = "custom_view_details_suggestion"
idempotent = true

[[inputs]]
name = "filter"
type = "JSONObject"
description = "The filter object to generate suggestions for."
required = true

[[inputs]]
name = "modelName"
type = "String"
description = "The entity type the view targets. If null, defaults to \"issue\". Valid values: \"issue\", \"project\", \"initiative\", \"feedItem\"."
required = false
+++
