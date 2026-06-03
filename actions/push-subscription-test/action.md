+++
name = "push-subscription-test"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/push-subscription-test@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["push_subscription_test"]

[match]
intent = "Sends a test push notification to the authenticated user's registered devices. Useful for verifying that push notification delivery is working correctly."

[[execute]]
id = "push"
connector = "github://ALRubinger/aileron-connector-linear"
op = "push_subscription_test"
idempotent = true

[[inputs]]
name = "sendStrategy"
type = "SendStrategy"
description = "The send strategy to use."
required = false

[[inputs]]
name = "targetMobile"
type = "Boolean"
description = "Whether to send to mobile devices."
required = false
+++
