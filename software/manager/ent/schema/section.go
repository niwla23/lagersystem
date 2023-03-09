package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/hook"
)

// Section holds the schema definition for the Section entity.
type Section struct {
	ent.Schema
}

// Fields of the Section.
func (Section) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
	}
}

// Edges of the Section.
func (Section) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("box", Box.Type).Ref("sections").Unique(),
		edge.From("parts", Part.Type).Ref("section"),
	}
}

// Hooks of the Section.
func (Section) Hooks() []ent.Hook {
	return []ent.Hook{
		// update updatedAt
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.SectionFunc(func(ctx context.Context, m *gen.SectionMutation) (ent.Value, error) {
					m.SetUpdatedAt(time.Now())
					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
