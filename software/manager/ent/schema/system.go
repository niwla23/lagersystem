package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// System holds the schema definition for the System entity.
type System struct {
	ent.Schema
}

// Fields of the System.
func (System) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.Time("modifiedAt").Default(time.Now),
		field.String("name").NotEmpty().Unique(),
		field.String("description"),
	}
}

// Edges of the System.
func (System) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("boxes", Box.Type),
	}
}
