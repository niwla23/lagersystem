// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BoxesColumns holds the columns for the "boxes" table.
	BoxesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// BoxesTable holds the schema information for the "boxes" table.
	BoxesTable = &schema.Table{
		Name:       "boxes",
		Columns:    BoxesColumns,
		PrimaryKey: []*schema.Column{BoxesColumns[0]},
	}
	// PartsColumns holds the columns for the "parts" table.
	PartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// PartsTable holds the schema information for the "parts" table.
	PartsTable = &schema.Table{
		Name:       "parts",
		Columns:    PartsColumns,
		PrimaryKey: []*schema.Column{PartsColumns[0]},
	}
	// PositionsColumns holds the columns for the "positions" table.
	PositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "position_id", Type: field.TypeInt, Unique: true},
		{Name: "box_position", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PositionsTable holds the schema information for the "positions" table.
	PositionsTable = &schema.Table{
		Name:       "positions",
		Columns:    PositionsColumns,
		PrimaryKey: []*schema.Column{PositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "positions_boxes_position",
				Columns:    []*schema.Column{PositionsColumns[3]},
				RefColumns: []*schema.Column{BoxesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PropertiesColumns holds the columns for the "properties" table.
	PropertiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "part_properties", Type: field.TypeInt, Nullable: true},
	}
	// PropertiesTable holds the schema information for the "properties" table.
	PropertiesTable = &schema.Table{
		Name:       "properties",
		Columns:    PropertiesColumns,
		PrimaryKey: []*schema.Column{PropertiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "properties_parts_properties",
				Columns:    []*schema.Column{PropertiesColumns[5]},
				RefColumns: []*schema.Column{PartsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SectionsColumns holds the columns for the "sections" table.
	SectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "box_sections", Type: field.TypeInt, Nullable: true},
	}
	// SectionsTable holds the schema information for the "sections" table.
	SectionsTable = &schema.Table{
		Name:       "sections",
		Columns:    SectionsColumns,
		PrimaryKey: []*schema.Column{SectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sections_boxes_sections",
				Columns:    []*schema.Column{SectionsColumns[2]},
				RefColumns: []*schema.Column{BoxesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "tag_children", Type: field.TypeInt, Nullable: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tags_tags_children",
				Columns:    []*schema.Column{TagsColumns[4]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PartTagsColumns holds the columns for the "part_tags" table.
	PartTagsColumns = []*schema.Column{
		{Name: "part_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// PartTagsTable holds the schema information for the "part_tags" table.
	PartTagsTable = &schema.Table{
		Name:       "part_tags",
		Columns:    PartTagsColumns,
		PrimaryKey: []*schema.Column{PartTagsColumns[0], PartTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "part_tags_part_id",
				Columns:    []*schema.Column{PartTagsColumns[0]},
				RefColumns: []*schema.Column{PartsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "part_tags_tag_id",
				Columns:    []*schema.Column{PartTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PartSectionsColumns holds the columns for the "part_sections" table.
	PartSectionsColumns = []*schema.Column{
		{Name: "part_id", Type: field.TypeInt},
		{Name: "section_id", Type: field.TypeInt},
	}
	// PartSectionsTable holds the schema information for the "part_sections" table.
	PartSectionsTable = &schema.Table{
		Name:       "part_sections",
		Columns:    PartSectionsColumns,
		PrimaryKey: []*schema.Column{PartSectionsColumns[0], PartSectionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "part_sections_part_id",
				Columns:    []*schema.Column{PartSectionsColumns[0]},
				RefColumns: []*schema.Column{PartsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "part_sections_section_id",
				Columns:    []*schema.Column{PartSectionsColumns[1]},
				RefColumns: []*schema.Column{SectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BoxesTable,
		PartsTable,
		PositionsTable,
		PropertiesTable,
		SectionsTable,
		TagsTable,
		PartTagsTable,
		PartSectionsTable,
	}
)

func init() {
	PositionsTable.ForeignKeys[0].RefTable = BoxesTable
	PropertiesTable.ForeignKeys[0].RefTable = PartsTable
	SectionsTable.ForeignKeys[0].RefTable = BoxesTable
	TagsTable.ForeignKeys[0].RefTable = TagsTable
	PartTagsTable.ForeignKeys[0].RefTable = PartsTable
	PartTagsTable.ForeignKeys[1].RefTable = TagsTable
	PartSectionsTable.ForeignKeys[0].RefTable = PartsTable
	PartSectionsTable.ForeignKeys[1].RefTable = SectionsTable
}
