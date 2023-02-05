package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Property holds the schema definition for the Property entity.
// Property represents a property of a part.
type Property struct {
	ent.Schema
}

// Fields of the Property.
func (Property) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.String("name").NotEmpty(),
		field.String("value").NotEmpty(),
		field.String("type").NotEmpty(),
	}
}

// Edges of the Property.
func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("part", Part.Type).Ref("properties").Unique(),
	}
}
