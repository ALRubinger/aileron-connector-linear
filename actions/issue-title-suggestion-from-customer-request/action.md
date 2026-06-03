+++
name = "issue-title-suggestion-from-customer-request"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-title-suggestion-from-customer-request@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_title_suggestion_from_customer_request"]

[match]
intent = "Suggests an issue title based on a customer request message using AI summarization."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_title_suggestion_from_customer_request"
idempotent = true

[[inputs]]
name = "request"
type = "String"
description = "The customer request message to generate an issue title from."
required = true
+++
