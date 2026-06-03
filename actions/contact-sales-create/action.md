+++
name = "contact-sales-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/contact-sales-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["contact_sales_create"]

[match]
intent = "[INTERNAL] Submits a sales pricing inquiry from the website. Small companies (1-19 employees) are routed to Intercom, while larger companies are routed to HubSpot."

[[execute]]
id = "contact"
connector = "github://ALRubinger/aileron-connector-linear"
op = "contact_sales_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "companySize"
type = "String"
description = "The size of the inquiring company (e.g., '1-19', '20-99', '100-499'). Used to route the inquiry to the appropriate sales channel."
required = false

[[inputs]]
name = "distinctId"
type = "String"
description = "PostHog distinct ID for correlating this inquiry with anonymous analytics events."
required = false

[[inputs]]
name = "email"
type = "String"
description = "Work email address of the person submitting the sales inquiry."
required = true

[[inputs]]
name = "message"
type = "String"
description = "An optional message from the user describing their needs or questions."
required = false

[[inputs]]
name = "name"
type = "String"
description = "Full name of the person submitting the sales inquiry."
required = true

[[inputs]]
name = "sessionId"
type = "String"
description = "PostHog session ID for correlating this inquiry with the user's browsing session."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The page URL from which the sales inquiry was submitted, for attribution tracking."
required = false

[[inputs]]
name = "utm"
type = "String"
description = "JSON serialized last-touch UTM parameters captured from the `utm` cookie at form submission."
required = false

[[inputs]]
name = "utmFirstTouch"
type = "String"
description = "JSON serialized first-touch UTM parameters captured from the `utm_first` cookie at form submission."
required = false
+++
