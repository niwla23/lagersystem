package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Part holds the schema definition for the Part entity.
type Part struct {
	ent.Schema
}

// Fields of the Part.
func (Part) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.String("name").NotEmpty().Unique(),
		field.String("description"),
	}
}

// Edges of the Part.
func (Part) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.To("properties", Property.Type),
		edge.To("sections", Section.Type),
	}
}
