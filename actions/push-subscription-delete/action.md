+++
name = "push-subscription-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/push-subscription-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["push_subscription_delete"]

[match]
intent = "Deletes a push subscription, unregistering the device from receiving push notifications."

[[execute]]
id = "push"
connector = "github://ALRubinger/aileron-connector-linear"
op = "push_subscription_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the push subscription to delete."
required = true
+++
