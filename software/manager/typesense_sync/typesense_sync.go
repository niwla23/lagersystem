package typesense_sync

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/niwla23/lagersystem/manager/typesense_wrapper"
	"github.com/typesense/typesense-go/typesense/api"

	ent "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/part"
)

var SYNC_INTERVAL = 30 * time.Second

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func syncPartsToTypesense(parts []*ent.Part) {
	// using interface type here might look like a dumb idea because we know the type
	// and it is but go is too incompetent to convert a specific type array to an interface array
	// so yeah that is how it is.
	var documentsToIndex []interface{}
	for _, part := range parts {
		partDocument := typesense_wrapper.PartDocument{
			ID:          strconv.Itoa(part.ID),
			Name:        part.Name,
			Description: part.Description,
		}
		documentsToIndex = append(documentsToIndex, partDocument)
	}
	if len(documentsToIndex) > 0 {
		fmt.Println("Indexing parts...")
		params := &api.ImportDocumentsParams{}

		_, err := typesense_wrapper.TypesenseClient.Collection("parts").Documents().Import(documentsToIndex, params)
		check(err)
	}
}

func SyncBackgroundTask() {
	for {
		fmt.Println("Syncing in background...")
		client, err := ent.Open("sqlite3", "file:///tmp/db.sqlite?_fk=1")
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
		ctx := context.Background()

		// get all non-deleted parts that have been modified since the last sync
		modifiedParts, err := client.Part.Query().Where(part.Deleted(false), part.UpdatedAtGT(time.Now().Add(SYNC_INTERVAL*-1))).All(ctx)
		check(err)
		fmt.Println("Modified parts:", modifiedParts)
		syncPartsToTypesense(modifiedParts)
		// todo: handle deleted parts

		// integrity check: check if part counts are equal
		partCount, err := client.Part.Query().Where(part.Deleted(false)).Count(ctx)
		check(err)
		resp, err := typesense_wrapper.TypesenseClient.Collection("parts").Retrieve()
		check(err)
		partCountTypesense := int(resp.NumDocuments)

		if partCount != partCountTypesense {
			fmt.Println("Integrity check failed, syncing all parts...")
			parts, err := client.Part.Query().Where(part.Deleted(false)).All(ctx)
			check(err)
			syncPartsToTypesense(parts)
		}

		time.Sleep(SYNC_INTERVAL)
	}
}
