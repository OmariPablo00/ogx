package feature

import "github.com/niconical/ogx/internal"

type Feature = internal.Flag

const (
	CTE Feature = 1 << iota
	WithValues
	Returning
	InsertReturning
	DefaultPlaceholder
	DoubleColonCast
	ValuesRow
	UpdateMultiTable
	InsertTableAlias
	UpdateTableAlias
	DeleteTableAlias
	AutoIncrement
	Identity
	TableCascade
	TableIdentity
	TableTruncate
	InsertOnConflict     // INSERT ... ON CONFLICT
	InsertOnDuplicateKey // INSERT ... ON DUPLICATE KEY
	InsertIgnore         // INSERT IGNORE ...
	TableNotExists
	OffsetFetch
	SelectExists
	UpdateFromTable
	MSSavepoint
	GeneratedIdentity
	CompositeIn // ... WHERE (A,B) IN ((N, NN), (N, NN)...)
)
