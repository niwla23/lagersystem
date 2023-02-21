package typesense_wrapper

import "github.com/typesense/typesense-go/typesense"

var TypesenseClient *typesense.Client

func InitTypesense() {
	TypesenseClient = typesense.NewClient(
		typesense.WithServer("http://localhost:8108"),
		typesense.WithAPIKey("xyz"))
}
