+++
name = "project-milestone-move"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-milestone-move@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_milestone_move"]

[match]
intent = "[Internal] Moves a project milestone to another project, can be called to undo a prior move."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_milestone_move"
idempotent = false

[approval]
required = true

[[inputs]]
name = "addIssueTeamToProject"
type = "Boolean"
description = "Whether to add each milestone issue's team to the project. This is needed when there is a mismatch between a project's teams and the milestone's issues' teams. Either this or newIssueTeamId is required in that situation to resolve constraints."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project milestone to move."
required = true

[[inputs]]
name = "newIssueTeamId"
type = "String"
description = "The team id to move the attached issues to. This is needed when there is a mismatch between a project's teams and the milestone's issues' teams. Either this or addIssueTeamToProject is required in that situation to resolve constraints."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "The identifier of the project to move the milestone to."
required = true

[[inputs]]
name = "undoIssueTeamIds"
type = "ProjectMilestoneMoveIssueToTeamInput"
description = "A list of issue id to team ids, used for undoing a previous milestone move where the specified issues were moved from the specified teams."
required = false

[[inputs]]
name = "undoProjectTeamIds"
type = "ProjectMilestoneMoveProjectTeamsInput"
description = "A mapping of project id to a previous set of team ids, used for undoing a previous milestone move where the specified teams were added to the project."
required = false
+++
