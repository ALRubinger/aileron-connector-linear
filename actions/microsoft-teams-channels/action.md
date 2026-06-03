+++
name = "microsoft-teams-channels"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/microsoft-teams-channels@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["microsoft_teams_channels"]

[match]
intent = "[Internal] Lists Microsoft Teams teams and channels accessible to the connecting user. Uses the user's microsoftPersonal integration for auth. Used by the project channel picker dialog to populate the team/channel dropdowns."

[[execute]]
id = "microsoft"
connector = "github://ALRubinger/aileron-connector-linear"
op = "microsoft_teams_channels"
idempotent = true
+++
