package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Box holds the schema definition for the Box entity.
type Box struct {
	ent.Schema
}

// Fields of the Box.
func (Box) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the Box.
func (Box) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sections", Section.Type),
		edge.To("position", Position.Type).
			Unique(),
	}
}
