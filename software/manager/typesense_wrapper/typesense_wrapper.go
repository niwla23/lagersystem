package typesense_wrapper

import (
	"fmt"

	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

var TypesenseClient *typesense.Client

type PartDocument struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func InitTypesense() {
	// create typesense client
	TypesenseClient = typesense.NewClient(
		typesense.WithServer("http://localhost:8108"),
		typesense.WithAPIKey("xyz"))

	// create typesense collection if it doesn't exist
	partCollection := TypesenseClient.Collection("parts")
	_, err := partCollection.Retrieve()
	if err != nil {
		if err.Error() == `status: 404 response: {"message": "Not Found"}` {
			schema := &api.CollectionSchema{
				Name: "parts",
				Fields: []api.Field{
					{
						Name: "id",
						Type: "string",
					},
					{
						Name: "name",
						Type: "string",
					},
					{
						Name: "description",
						Type: "string",
					},
					{
						Name: "tags",
						Type: "string[]",
					},
				},
			}

			_, err := TypesenseClient.Collections().Create(schema)
			if err != nil {
				panic(fmt.Sprintf("failed creating typesense collection: %v", err))
			}
		} else {
			panic(fmt.Sprintf("failed getting typesense collection: %v", err))
		}
	}
}
