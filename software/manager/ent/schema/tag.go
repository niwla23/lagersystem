package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/niwla23/lagersystem/manager/ent"
	"github.com/niwla23/lagersystem/manager/ent/hook"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
		field.String("name").NotEmpty().Unique(),
		field.String("description"),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parts", Part.Type).Ref("tags"),
		edge.To("children", Tag.Type).From("parent").Unique(),
	}
}

// Hooks of the Tag.
func (Tag) Hooks() []ent.Hook {
	return []ent.Hook{
		// update updatedAt
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.TagFunc(func(ctx context.Context, m *gen.TagMutation) (ent.Value, error) {
					m.SetUpdatedAt(time.Now())
					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
