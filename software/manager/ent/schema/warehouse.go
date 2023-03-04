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

// Warehouse holds the schema definition for the Warehouse entity.
type Warehouse struct {
	ent.Schema
}

// Fields of the Warehouse.
func (Warehouse) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
		field.String("name").NotEmpty().Unique(),
		field.String("description"),
	}
}

// Edges of the Warehouse.
func (Warehouse) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("positions", Position.Type),
	}
}

// Hooks of the Warehouse.
func (Warehouse) Hooks() []ent.Hook {
	return []ent.Hook{
		// update updatedAt
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.WarehouseFunc(func(ctx context.Context, m *gen.WarehouseMutation) (ent.Value, error) {
					m.SetUpdatedAt(time.Now())
					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
