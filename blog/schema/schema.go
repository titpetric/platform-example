package schema

import _ "embed"

// InitialSchema contains the initial blog schema
// This is executed by the storage package on first use
//
//go:embed 2025-01-01-000000-articles-initial.up.sql
var InitialSchema string
