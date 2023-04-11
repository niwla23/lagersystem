package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	gen "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/hook"
)

// Box holds the schema definition for the Box entity.
type Box struct {
	ent.Schema
}

// Fields of the Box.
func (Box) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
		field.UUID("boxId", uuid.UUID{}).Unique(),
	}
}

// Edges of the Box.
func (Box) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parts", Part.Type),
		edge.To("position", Position.Type).Unique().StructTag(`json:"position"`),
	}
}

// Hooks of the Box.
func (Box) Hooks() []ent.Hook {
	return []ent.Hook{
		// update updatedAt
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.BoxFunc(func(ctx context.Context, m *gen.BoxMutation) (ent.Value, error) {
					m.SetUpdatedAt(time.Now())
					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
