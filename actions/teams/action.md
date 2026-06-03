+++
name = "teams"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/teams@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["teams"]

[match]
intent = "All teams whose issues the user can access. This includes public teams and private teams the user is a member of. This may differ from `administrableTeams`, which returns teams whose settings the user can change but whose issues they don't necessarily have access to."

[[execute]]
id = "teams"
connector = "github://ALRubinger/aileron-connector-linear"
op = "teams"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "ancestors"
type = "TeamCollectionFilter"
description = "Filters that the team's ancestors must satisfy."
required = false

[[inputs]]
name = "and"
type = "TeamFilter"
description = "Compound filters, all of which need to be matched by the team."
required = false

[[inputs]]
name = "before"
type = "String"
description = "A cursor to be used with last for backward pagination."
required = false

[[inputs]]
name = "createdAt"
type = "DateComparator"
description = "Comparator for the created at date."
required = false

[[inputs]]
name = "description"
type = "NullableStringComparator"
description = "Comparator for the team description."
required = false

[[inputs]]
name = "first"
type = "Int"
description = "The number of items to forward paginate (used with after). Defaults to 50."
required = false

[[inputs]]
name = "id"
type = "IDComparator"
description = "Comparator for the identifier."
required = false

[[inputs]]
name = "includeArchived"
type = "Boolean"
description = "Should archived resources be included (default: false)"
required = false

[[inputs]]
name = "issues"
type = "IssueCollectionFilter"
description = "Filters that the teams issues must satisfy."
required = false

[[inputs]]
name = "key"
type = "StringComparator"
description = "Comparator for the team key."
required = false

[[inputs]]
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "members"
type = "UserCollectionFilter"
description = "Filters that the team's members must satisfy."
required = false

[[inputs]]
name = "name"
type = "StringComparator"
description = "Comparator for the team name."
required = false

[[inputs]]
name = "or"
type = "TeamFilter"
description = "Compound filters, one of which need to be matched by the team."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "parent"
type = "NullableTeamFilter"
description = "Filters that the teams parent must satisfy."
required = false

[[inputs]]
name = "private"
type = "BooleanComparator"
description = "Comparator for the team privacy."
required = false

[[inputs]]
name = "releasePipelines"
type = "ReleasePipelineCollectionFilter"
description = "Filters that the team's release pipelines must satisfy."
required = false

[[inputs]]
name = "retiredAt"
type = "NullableDateComparator"
description = "Comparator for the time at which the team was retired."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false
+++
