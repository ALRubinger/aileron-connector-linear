+++
name = "attachment-link-git-lab-m-r"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-link-git-lab-m-r@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_link_git_lab_m_r"]

[match]
intent = "Link an existing GitLab MR to an issue. This creates a rich attachment using the workspace's GitLab integration, enabling features like automated status syncing."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_link_git_lab_m_r"
idempotent = false

[approval]
required = true

[[inputs]]
name = "createAsUser"
type = "String"
description = "Create attachment as a user with the provided name. This option is only available to OAuth applications creating attachments in `actor=app` mode."
required = false

[[inputs]]
name = "displayIconUrl"
type = "String"
description = "Provide an external user avatar URL. Can only be used in conjunction with the `createAsUser` options. This option is only available to OAuth applications creating comments in `actor=app` mode."
required = false

[[inputs]]
name = "id"
type = "String"
description = "Optional attachment ID that may be provided through the API."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The issue for which to link the GitLab merge request. Can be a UUID or issue identifier (e.g., 'LIN-123')."
required = true

[[inputs]]
name = "number"
type = "Float"
description = "The GitLab merge request number to link."
required = true

[[inputs]]
name = "projectPathWithNamespace"
type = "String"
description = "The path name to the project including any (sub)groups. E.g. linear/main/client."
required = true

[[inputs]]
name = "title"
type = "String"
description = "The title to use for the attachment."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The URL of the GitLab merge request to link."
required = true
+++
