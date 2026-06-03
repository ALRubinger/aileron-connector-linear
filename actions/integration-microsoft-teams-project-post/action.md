+++
name = "integration-microsoft-teams-project-post"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-microsoft-teams-project-post@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_microsoft_teams_project_post"]

[match]
intent = "[Internal] Connect a project to a Microsoft Teams channel. Find-or-update semantics: creates a microsoftTeamsProjectPost integration row if none exists for the project, or overwrites the existing one's team/channel selection. Requires the connecting user to have linked their personal Microsoft account."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_microsoft_teams_project_post"
idempotent = false

[approval]
required = true

[[inputs]]
name = "channelId"
type = "String"
description = "Microsoft Teams channel id."
required = true

[[inputs]]
name = "channelName"
type = "String"
description = "Display name of the chosen channel."
required = true

[[inputs]]
name = "membershipType"
type = "String"
description = "Membership type of the channel: standard, private, or shared."
required = true

[[inputs]]
name = "projectId"
type = "String"
description = "Integration's associated project."
required = true

[[inputs]]
name = "teamId"
type = "String"
description = "AAD group id of the chosen team."
required = true

[[inputs]]
name = "teamName"
type = "String"
description = "Display name of the chosen team."
required = true
+++
