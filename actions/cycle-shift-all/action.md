+++
name = "cycle-shift-all"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/cycle-shift-all@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["cycle_shift_all"]

[match]
intent = "Shifts all cycles starts and ends by a certain number of days, starting from the provided cycle onwards."

[[execute]]
id = "cycle"
connector = "github://ALRubinger/aileron-connector-linear"
op = "cycle_shift_all"
idempotent = false

[approval]
required = true

[[inputs]]
name = "daysToShift"
type = "Float"
description = "The number of days to shift the cycles by."
required = true

[[inputs]]
name = "id"
type = "String"
description = "The cycle ID at which to start the shift."
required = true
+++
