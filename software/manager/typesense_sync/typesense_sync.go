package typesense_sync

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/niwla23/lagersystem/manager/config"
	"github.com/niwla23/lagersystem/manager/typesense_wrapper"
	"github.com/typesense/typesense-go/typesense/api"

	ent "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/part"
)

var SYNC_INTERVAL = 10 * time.Second

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

		if part.Deleted {
			_, err := typesense_wrapper.TypesenseClient.Collection("parts").Document(strconv.Itoa(part.ID)).Delete()
			check(err)
			continue
		}

		tags := []string{}
		for _, tag := range part.Edges.Tags {
			tags = append(tags, tag.Name)
		}
		fmt.Println(tags)

		partDocument := typesense_wrapper.PartDocument{
			ID:          strconv.Itoa(part.ID),
			Name:        part.Name,
			Description: part.Description,
			Tags:        tags,
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
		client, err := ent.Open("sqlite3", config.DBUri)
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
		ctx := context.Background()

		// get all parts that have been modified since the last sync
		modifiedParts, err := client.Part.Query().Where(part.UpdatedAtGT(time.Now().Add(SYNC_INTERVAL * -1))).WithTags().All(ctx)
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
			fmt.Println("Integrity check failed, deleting collection and syncing all parts...")
			fmt.Println("Part count DB:", partCount, "Part count Typesense:", partCountTypesense)
			parts, err := client.Part.Query().Where(part.Deleted(false)).WithTags().All(ctx)
			check(err)
			_, err = typesense_wrapper.TypesenseClient.Collection("parts").Delete()
			check(err)
			typesense_wrapper.InitTypesense()
			syncPartsToTypesense(parts)
		}

		time.Sleep(SYNC_INTERVAL)
	}
}
