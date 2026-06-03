+++
name = "dummy"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/dummy@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["dummy"]

[match]
intent = ""

[[execute]]
id = "dummy"
connector = "github://ALRubinger/aileron-connector-linear"
op = "dummy"
idempotent = true
+++
