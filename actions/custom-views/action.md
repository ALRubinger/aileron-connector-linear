+++
name = "custom-views"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/custom-views@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["custom_views"]

[match]
intent = "All custom views accessible to the user, including personal views and shared workspace views. Excludes views scoped to a specific project or initiative."

[[execute]]
id = "custom"
connector = "github://ALRubinger/aileron-connector-linear"
op = "custom_views"
idempotent = true

[[inputs]]
name = "after"
type = "String"
description = "A cursor to be used with first for forward pagination"
required = false

[[inputs]]
name = "and"
type = "CustomViewFilter"
description = "Compound filters, all of which need to be matched by the custom view."
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
name = "createdAt"
type = "CustomViewCreatedAtSort"
description = "Sort by custom view creation date."
required = false

[[inputs]]
name = "creator"
type = "UserFilter"
description = "Filters that the custom view creator must satisfy."
required = false

[[inputs]]
name = "first"
type = "Int"
description = "The number of items to forward paginate (used with after). Defaults to 50."
required = false

[[inputs]]
name = "hasFacet"
type = "Boolean"
description = "[INTERNAL] Filter based on whether the custom view has a facet."
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
name = "last"
type = "Int"
description = "The number of items to backward paginate (used with before). Defaults to 50."
required = false

[[inputs]]
name = "modelName"
type = "StringComparator"
description = "Comparator for the custom view model name."
required = false

[[inputs]]
name = "name"
type = "StringComparator"
description = "Comparator for the custom view name."
required = false

[[inputs]]
name = "name"
type = "CustomViewNameSort"
description = "Sort by custom view name."
required = false

[[inputs]]
name = "or"
type = "CustomViewFilter"
description = "Compound filters, one of which need to be matched by the custom view."
required = false

[[inputs]]
name = "orderBy"
type = "PaginationOrderBy"
description = "By which field should the pagination order by. Available options are createdAt (default) and updatedAt."
required = false

[[inputs]]
name = "shared"
type = "BooleanComparator"
description = "Comparator for whether the custom view is shared."
required = false

[[inputs]]
name = "shared"
type = "CustomViewSharedSort"
description = "Sort by custom view shared status."
required = false

[[inputs]]
name = "team"
type = "NullableTeamFilter"
description = "Filters that the custom view's team must satisfy."
required = false

[[inputs]]
name = "updatedAt"
type = "DateComparator"
description = "Comparator for the updated at date."
required = false

[[inputs]]
name = "updatedAt"
type = "CustomViewUpdatedAtSort"
description = "Sort by custom view update date."
required = false
+++
