+++
name = "audit-entry-types"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/audit-entry-types@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["audit_entry_types"]

[match]
intent = "List of audit entry types."

[[execute]]
id = "audit"
connector = "github://ALRubinger/aileron-connector-linear"
op = "audit_entry_types"
idempotent = true
+++
