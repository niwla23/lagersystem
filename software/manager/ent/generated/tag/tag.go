// Code generated by ent, DO NOT EDIT.

package tag

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeParts holds the string denoting the parts edge name in mutations.
	EdgeParts = "parts"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// PartsTable is the table that holds the parts relation/edge. The primary key declared below.
	PartsTable = "part_tags"
	// PartsInverseTable is the table name for the Part entity.
	// It exists in this package in order to avoid circular dependency with the "part" package.
	PartsInverseTable = "parts"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "tags"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "tag_children"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "tags"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "tag_children"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tags"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"tag_children",
}

var (
	// PartsPrimaryKey and PartsColumn2 are the table columns denoting the
	// primary key for the parts relation (M2M).
	PartsPrimaryKey = []string{"part_id", "tag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)