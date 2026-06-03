+++
name = "integration-settings-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-settings-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_settings_update"]

[match]
intent = "[INTERNAL] Updates the integration settings."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_settings_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "front"
type = "FrontSettingsInput"
description = ""
required = false

[[inputs]]
name = "gitHub"
type = "GitHubSettingsInput"
description = ""
required = false

[[inputs]]
name = "gitHubImport"
type = "GitHubImportSettingsInput"
description = ""
required = false

[[inputs]]
name = "gitHubPersonal"
type = "GitHubPersonalSettingsInput"
description = ""
required = false

[[inputs]]
name = "gitLab"
type = "GitLabSettingsInput"
description = ""
required = false

[[inputs]]
name = "gong"
type = "GongSettingsInput"
description = ""
required = false

[[inputs]]
name = "googleSheets"
type = "GoogleSheetsSettingsInput"
description = ""
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the integration to update."
required = true

[[inputs]]
name = "intercom"
type = "IntercomSettingsInput"
description = ""
required = false

[[inputs]]
name = "jira"
type = "JiraSettingsInput"
description = ""
required = false

[[inputs]]
name = "jiraPersonal"
type = "JiraPersonalSettingsInput"
description = ""
required = false

[[inputs]]
name = "launchDarkly"
type = "LaunchDarklySettingsInput"
description = ""
required = false

[[inputs]]
name = "microsoftTeams"
type = "MicrosoftTeamsSettingsInput"
description = ""
required = false

[[inputs]]
name = "microsoftTeamsProjectPost"
type = "MicrosoftTeamsPostSettingsInput"
description = ""
required = false

[[inputs]]
name = "notion"
type = "NotionSettingsInput"
description = ""
required = false

[[inputs]]
name = "opsgenie"
type = "OpsgenieInput"
description = ""
required = false

[[inputs]]
name = "pagerDuty"
type = "PagerDutyInput"
description = ""
required = false

[[inputs]]
name = "salesforce"
type = "SalesforceSettingsInput"
description = ""
required = false

[[inputs]]
name = "sentry"
type = "SentrySettingsInput"
description = ""
required = false

[[inputs]]
name = "slack"
type = "SlackSettingsInput"
description = ""
required = false

[[inputs]]
name = "slackAsks"
type = "SlackAsksSettingsInput"
description = ""
required = false

[[inputs]]
name = "slackCustomViewNotifications"
type = "SlackPostSettingsInput"
description = ""
required = false

[[inputs]]
name = "slackInitiativePost"
type = "SlackPostSettingsInput"
description = ""
required = false

[[inputs]]
name = "slackOrgInitiativeUpdatesPost"
type = "SlackPostSettingsInput"
description = ""
required = false

[[inputs]]
name = "slackOrgProjectUpdatesPost"
type = "SlackPostSettingsInput"
description = ""
required = false

[[inputs]]
name = "slackPost"
type = "SlackPostSettingsInput"
description = ""
required = false

[[inputs]]
name = "slackProjectPost"
type = "SlackPostSettingsInput"
description = ""
required = false

[[inputs]]
name = "zendesk"
type = "ZendeskSettingsInput"
description = ""
required = false
+++
