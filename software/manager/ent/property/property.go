// Code generated by ent, DO NOT EDIT.

package property

import (
	"time"
)

const (
	// Label holds the string label denoting the property type in the database.
	Label = "property"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgePart holds the string denoting the part edge name in mutations.
	EdgePart = "part"
	// Table holds the table name of the property in the database.
	Table = "properties"
	// PartTable is the table that holds the part relation/edge.
	PartTable = "properties"
	// PartInverseTable is the table name for the Part entity.
	// It exists in this package in order to avoid circular dependency with the "part" package.
	PartInverseTable = "parts"
	// PartColumn is the table column denoting the part relation/edge.
	PartColumn = "part_properties"
)

// Columns holds all SQL columns for property fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldName,
	FieldValue,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "properties"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"part_properties",
}

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

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
)
