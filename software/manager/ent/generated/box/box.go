// Code generated by ent, DO NOT EDIT.

package box

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the box type in the database.
	Label = "box"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldBoxId holds the string denoting the boxid field in the database.
	FieldBoxId = "box_id"
	// EdgeParts holds the string denoting the parts edge name in mutations.
	EdgeParts = "parts"
	// EdgePosition holds the string denoting the position edge name in mutations.
	EdgePosition = "position"
	// Table holds the table name of the box in the database.
	Table = "boxes"
	// PartsTable is the table that holds the parts relation/edge.
	PartsTable = "parts"
	// PartsInverseTable is the table name for the Part entity.
	// It exists in this package in order to avoid circular dependency with the "part" package.
	PartsInverseTable = "parts"
	// PartsColumn is the table column denoting the parts relation/edge.
	PartsColumn = "box_parts"
	// PositionTable is the table that holds the position relation/edge.
	PositionTable = "positions"
	// PositionInverseTable is the table name for the Position entity.
	// It exists in this package in order to avoid circular dependency with the "position" package.
	PositionInverseTable = "positions"
	// PositionColumn is the table column denoting the position relation/edge.
	PositionColumn = "box_position"
)

// Columns holds all SQL columns for box fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldBoxId,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/niwla23/lagersystem/manager/ent/generated/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
)
