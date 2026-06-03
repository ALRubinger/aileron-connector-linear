+++
name = "user-unlink-from-identity-provider"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-unlink-from-identity-provider@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_unlink_from_identity_provider"]

[match]
intent = "Unlinks a guest user from their identity provider. Can only be called by an admin when SCIM is enabled."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_unlink_from_identity_provider"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the guest user to unlink from their identity provider."
required = true
+++
