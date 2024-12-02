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

// Position holds the schema definition for the Position entity.
type Position struct {
	ent.Schema
}

// Fields of the Position.
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
	}
}

// Edges of the Position.
func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("storedBox", Box.Type).Ref("position").Unique(),
		edge.From("warehouse", Warehouse.Type).Ref("positions").Unique(),
	}
}

// Hooks of the Position.
func (Position) Hooks() []ent.Hook {
	return []ent.Hook{
		// update updatedAt
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.PositionFunc(func(ctx context.Context, m *gen.PositionMutation) (ent.Value, error) {
					m.SetUpdatedAt(time.Now())
					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
