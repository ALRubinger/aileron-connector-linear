+++
name = "rate-limit-status"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/rate-limit-status@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["rate_limit_status"]

[match]
intent = "The current rate limit status for the authenticated client, including remaining quota and reset timing for each limit type."

[[execute]]
id = "rate"
connector = "github://ALRubinger/aileron-connector-linear"
op = "rate_limit_status"
idempotent = true
+++
