+++
name = "initiative-update-mutation"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-update-mutation@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_update_mutation"]

[match]
intent = "Updates an initiative."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_update_mutation"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The initiative's color."
required = false

[[inputs]]
name = "content"
type = "String"
description = "The initiative's content in markdown format."
required = false

[[inputs]]
name = "customIdentifier"
type = "String"
description = "[Internal] Sets or clears the custom identifier override for the initiative. Pass null to revert to the workspace default."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the initiative."
required = false

[[inputs]]
name = "frequencyResolution"
type = "FrequencyResolutionType"
description = "The resolution type for the update reminder frequency (e.g., weekly, biweekly)."
required = false

[[inputs]]
name = "icon"
type = "String"
description = "The initiative's icon."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the initiative to update."
required = true

[[inputs]]
name = "labelIds"
type = "String"
description = "[Internal] The identifiers of the initiative labels associated with this initiative."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the initiative."
required = false

[[inputs]]
name = "ownerId"
type = "String"
description = "The owner of the initiative."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order of the initiative within the workspace."
required = false

[[inputs]]
name = "status"
type = "InitiativeStatus"
description = "The initiative's status."
required = false

[[inputs]]
name = "targetDate"
type = "TimelessDate"
description = "The estimated completion date of the initiative. Set to null to clear."
required = false

[[inputs]]
name = "targetDateResolution"
type = "DateResolutionType"
description = "The resolution of the initiative's estimated completion date."
required = false

[[inputs]]
name = "trashed"
type = "Boolean"
description = "Whether the initiative has been trashed. Set to true to trash, or null to restore."
required = false

[[inputs]]
name = "updateReminderFrequency"
type = "Float"
description = "The frequency at which to prompt for initiative updates. When not set, reminders are inherited from workspace settings."
required = false

[[inputs]]
name = "updateReminderFrequencyInWeeks"
type = "Float"
description = "The n-weekly frequency at which to prompt for initiative updates. When not set, reminders are inherited from workspace settings."
required = false

[[inputs]]
name = "updateRemindersDay"
type = "Day"
description = "The day of the week on which to prompt for initiative updates."
required = false

[[inputs]]
name = "updateRemindersHour"
type = "Int"
description = "The hour of the day (0-23) at which to prompt for initiative updates."
required = false
+++
