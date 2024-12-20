package typesense_sync

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/niwla23/lagersystem/manager/config"
	"github.com/niwla23/lagersystem/manager/database"
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

func SyncPartsToTypesense(parts []*ent.Part) {
	// using interface type here might look like a dumb idea because we know the type
	// and it is but go is too incompetent to convert a specific type array to an interface array
	// so yeah that is how it is.
	var documentsToIndex []interface{}
	for _, part := range parts {

		if part.Deleted {
			_, err := typesense_wrapper.TypesenseClient.Collection("parts").Document(strconv.Itoa(part.ID)).Delete()
			database.Client.Part.DeleteOne(part).ExecX(context.Background())
			check(err)
			continue
		}

		tags := []string{}
		for _, tag := range part.Edges.Tags {
			tags = append(tags, tag.Name)
		}

		partDocument := typesense_wrapper.PartDocument{
			ID:          strconv.Itoa(part.ID),
			Name:        part.Name,
			Description: part.Description,
			Tags:        tags,
			HasBox:      part.Edges.Box != nil,
			IsStored:    part.Edges.Box != nil && part.Edges.Box.Edges.Position != nil,
		}
		documentsToIndex = append(documentsToIndex, partDocument)
	}
	if len(documentsToIndex) > 0 {
		fmt.Println("Indexing parts...")
		action := "upsert"
		params := &api.ImportDocumentsParams{Action: &action}

		_, err := typesense_wrapper.TypesenseClient.Collection("parts").Documents().Import(documentsToIndex, params)

		check(err)
	}
}

func SyncBackgroundTask() {
	for {
		client, err := ent.Open("sqlite3", config.DBUri)
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
		ctx := context.Background()

		// get all parts that have been modified since the last sync
		modifiedParts, err := client.Part.Query().Where(part.UpdatedAtGT(time.Now().Add(SYNC_INTERVAL * -1))).WithTags().WithBox(func(q *ent.BoxQuery) { q.WithPosition() }).All(ctx)
		check(err)
		SyncPartsToTypesense(modifiedParts)

		// integrity check: check if part counts are equal
		partCount, err := client.Part.Query().Where(part.Deleted(false)).Count(ctx)
		check(err)
		resp, err := typesense_wrapper.TypesenseClient.Collection("parts").Retrieve()
		check(err)
		partCountTypesense := int(resp.NumDocuments)

		if partCount != partCountTypesense {
			fmt.Println("Integrity check failed, deleting collection and syncing all parts...")
			fmt.Println("Part count DB:", partCount, "Part count Typesense:", partCountTypesense)
			parts, err := client.Part.Query().Where(part.Deleted(false)).WithTags().WithBox(func(q *ent.BoxQuery) { q.WithPosition() }).All(ctx)
			check(err)
			_, err = typesense_wrapper.TypesenseClient.Collection("parts").Delete()
			check(err)
			typesense_wrapper.InitTypesense()
			SyncPartsToTypesense(parts)
		}

		time.Sleep(SYNC_INTERVAL)
	}
}
