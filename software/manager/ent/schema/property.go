package schema

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/hook"
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
		field.Time("updatedAt").Default(time.Now),
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

// Hooks of the Property.
func (Property) Hooks() []ent.Hook {
	return []ent.Hook{
		// First hook.
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.PropertyFunc(func(ctx context.Context, m *gen.PropertyMutation) (ent.Value, error) {
					propertyType, _ := m.GetType()
					propertyValue, _ := m.Value()
					switch propertyType {
					case "string":
						_, err := strconv.Atoi(propertyValue)
						if err == nil {
							return nil, fmt.Errorf("property value is a number but type is string")
						}
					case "number":
						_, err := strconv.Atoi(propertyValue)
						if err != nil {
							return nil, fmt.Errorf("cant convert property value to number")
						}
					case "boolean":
						if propertyValue != "true" && propertyValue != "false" {
							return nil, fmt.Errorf("property value is not a boolean")
						}
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.PropertyFunc(func(ctx context.Context, m *gen.PropertyMutation) (ent.Value, error) {
					m.SetUpdatedAt(time.Now())
					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
