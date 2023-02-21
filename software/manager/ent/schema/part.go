package schema

import (
	"strconv"
	"time"

	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	gen "github.com/niwla23/lagersystem/manager/ent"
	"github.com/niwla23/lagersystem/manager/ent/hook"
	"github.com/niwla23/lagersystem/manager/typesense_wrapper"
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

// Hooks of the Card.
func (Part) Hooks() []ent.Hook {
	return []ent.Hook{
		// First hook.
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.PartFunc(func(ctx context.Context, m *gen.PartMutation) (ent.Value, error) {
					fmt.Println("hook called")
					fmt.Println(m.Name())

					partId, _ := m.ID()
					fmt.Println("hook called1")

					part, err := m.Client().Part.Get(ctx, partId)
					if err != nil {
						fmt.Println("error mutating", err, partId)
						return next.Mutate(ctx, m)
					}
					fmt.Println(part.Name)

					document := struct {
						ID          string `json:"id"`
						Name        string `json:"name"`
						Description string `json:"description"`
					}{
						ID:          strconv.Itoa(part.ID),
						Name:        part.Name,
						Description: part.Description,
					}
					fmt.Println("hook called3")

					typesense_wrapper.TypesenseClient.Collection("parts").Documents().Upsert(document)
					fmt.Println("hook called4")

					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.PartFunc(func(ctx context.Context, m *gen.PartMutation) (ent.Value, error) {

					partId, _ := m.ID()

					typesense_wrapper.TypesenseClient.Collection("parts").Document(strconv.Itoa(partId)).Delete()

					return next.Mutate(ctx, m)
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}
