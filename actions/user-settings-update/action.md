+++
name = "user-settings-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-settings-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_settings_update"]

[match]
intent = "Updates the authenticated user's settings, including notification preferences, email subscriptions, theme, and other UI preferences."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_settings_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "feedLastSeenTime"
type = "DateTime"
description = "[Internal] The user's last seen time for the pulse feed."
required = false

[[inputs]]
name = "feedSummarySchedule"
type = "FeedSummarySchedule"
description = "[Internal] How often to generate a feed summary."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the userSettings to update."
required = true

[[inputs]]
name = "notificationCategoryPreferences"
type = "NotificationCategoryPreferencesInput"
description = "The user's notification category preferences."
required = false

[[inputs]]
name = "notificationChannelPreferences"
type = "PartialNotificationChannelPreferencesInput"
description = "The user's notification channel preferences."
required = false

[[inputs]]
name = "notificationDeliveryPreferences"
type = "NotificationDeliveryPreferencesInput"
description = "The user's notification delivery preferences."
required = false

[[inputs]]
name = "settings"
type = "JSONObject"
description = "The user's settings."
required = false

[[inputs]]
name = "subscribedToChangelog"
type = "Boolean"
description = "Whether this user is subscribed to changelog email or not."
required = false

[[inputs]]
name = "subscribedToDPA"
type = "Boolean"
description = "Whether this user is subscribed to DPA emails or not."
required = false

[[inputs]]
name = "subscribedToGeneralMarketingCommunications"
type = "Boolean"
description = "Whether this user is subscribed to general marketing communications or not."
required = false

[[inputs]]
name = "subscribedToInviteAccepted"
type = "Boolean"
description = "Whether this user is subscribed to invite accepted emails or not."
required = false

[[inputs]]
name = "subscribedToPrivacyLegalUpdates"
type = "Boolean"
description = "Whether this user is subscribed to privacy and legal update emails or not."
required = false

[[inputs]]
name = "usageWarningHistory"
type = "JSONObject"
description = "[Internal] The user's usage warning history."
required = false
+++
