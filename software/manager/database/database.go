package database

import (
	"context"
	"log"

	"github.com/niwla23/lagersystem/manager/config"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"
)

var Client *ent.Client

func InitDB() {
	client, err := ent.Open("sqlite3", config.DBUri)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// make client globally available
	Client = client
}
