// Code generated by ent, DO NOT EDIT.

package part

import (
	"time"
)

const (
	// Label holds the string label denoting the part type in the database.
	Label = "part"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeProperties holds the string denoting the properties edge name in mutations.
	EdgeProperties = "properties"
	// EdgeSections holds the string denoting the sections edge name in mutations.
	EdgeSections = "sections"
	// Table holds the table name of the part in the database.
	Table = "parts"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "part_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// PropertiesTable is the table that holds the properties relation/edge.
	PropertiesTable = "properties"
	// PropertiesInverseTable is the table name for the Property entity.
	// It exists in this package in order to avoid circular dependency with the "property" package.
	PropertiesInverseTable = "properties"
	// PropertiesColumn is the table column denoting the properties relation/edge.
	PropertiesColumn = "part_properties"
	// SectionsTable is the table that holds the sections relation/edge. The primary key declared below.
	SectionsTable = "part_sections"
	// SectionsInverseTable is the table name for the Section entity.
	// It exists in this package in order to avoid circular dependency with the "section" package.
	SectionsInverseTable = "sections"
)

// Columns holds all SQL columns for part fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldName,
	FieldDescription,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"part_id", "tag_id"}
	// SectionsPrimaryKey and SectionsColumn2 are the table columns denoting the
	// primary key for the sections relation (M2M).
	SectionsPrimaryKey = []string{"part_id", "section_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
)
