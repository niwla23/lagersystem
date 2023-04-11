package handlers

import (
	"context"
	"errors"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/niwla23/lagersystem/manager/config"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/part"
	"github.com/niwla23/lagersystem/manager/ent/generated/property"
	"github.com/niwla23/lagersystem/manager/ent/generated/tag"
	"github.com/niwla23/lagersystem/manager/helpers"
	"github.com/niwla23/lagersystem/manager/typesense_wrapper"
	"github.com/typesense/typesense-go/typesense/api"
)

type PropertyAddData struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type PartAddData struct {
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Amount      *int                       `json:"amount,omitempty"`
	Tags        []string                   `json:"tags"`
	Properties  map[string]PropertyAddData `json:"properties"`
	BoxId       int                        `json:"boxId"`
}

func createOrGetTagsFromNameList(tagNames *[]string, client *ent.Client, ctx context.Context) ([]*ent.Tag, error) {
	tags := make([]*ent.Tag, 0)
	for _, tagName := range *tagNames {
		var tagX *ent.Tag

		// try creating tag
		tagX, err := client.Tag.Create().SetName(tagName).SetDescription("").Save(ctx)

		if err != nil {
			// fetch tag from db
			target := &ent.ConstraintError{}
			if errors.As(err, &target) {
				tagX, err = client.Tag.Query().Where(tag.Name(tagName)).Only(ctx)
				if err != nil {
					return nil, err
				}
			} else {
				return nil, err
			}
		}
		tags = append(tags, tagX)
	}
	return tags, nil
}

func createOrUpdatePropertiesFromMap(part *ent.Part, properties *map[string]PropertyAddData, client *ent.Client, ctx context.Context) error {
	// delete unused properties
	for _, propertyX := range part.Edges.Properties {
		if _, ok := (*properties)[propertyX.Name]; !ok {
			err := client.Property.DeleteOne(propertyX).Exec(ctx)
			if err != nil {
				return err
			}
		}
	}

	for key, propertyData := range *properties {
		propertyX, err := part.QueryProperties().Where(property.Name(key)).Only(ctx)
		if err != nil {
			// property does not exist, create it
			_, err := client.Property.Create().SetName(key).SetValue(propertyData.Value).SetType(propertyData.Type).SetPart(part).Save(ctx)
			if err != nil {
				return err
			}
			continue
		}
		// property value or type is different, update it
		if propertyX.Value != propertyData.Value || propertyX.Type != propertyData.Type {
			_, err = propertyX.Update().SetValue(propertyData.Value).SetType(propertyData.Type).Save(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RegisterPartRoutes(router fiber.Router, client *ent.Client, ctx context.Context) {
	router.Post("/", func(c *fiber.Ctx) error {
		data := new(PartAddData)

		if err := c.BodyParser(data); err != nil {
			return err
		}

		// set default amount of -1 (unknown) unless the amount is given
		amount := -1
		if data.Amount != nil {
			amount = *data.Amount
		}

		// get tags from db or create them
		tags, err := createOrGetTagsFromNameList(&data.Tags, client, ctx)
		if err != nil {
			return err
		}

		// get box by given ID
		boxX, err := client.Box.Get(ctx, data.BoxId)
		if err != nil {
			return err
		}

		partX, err := client.Part.Create().
			SetName(data.Name).
			SetDescription(data.Description).
			SetAmount(amount).
			SetBox(boxX).
			AddTags(tags...).
			Save(ctx)

		if err != nil {
			return err
		}

		// add properties
		err = createOrUpdatePropertiesFromMap(partX, &data.Properties, client, ctx)
		if err != nil {
			return err
		}

		return c.JSON(partX)
	})

	router.Put("/:partId<int>", func(c *fiber.Ctx) error {
		partId, _ := strconv.Atoi(c.Params("partId"))

		// parse request body
		data := new(PartAddData)
		if err := c.BodyParser(data); err != nil {
			return err
		}

		// set default amount of -1 (unknown) unless the amount is given
		amount := -1
		if data.Amount != nil {
			amount = *data.Amount
		}

		// get partX from db
		partX, err := client.Part.Query().Where(part.ID(partId)).WithTags().WithProperties().Only(ctx)
		if err != nil {
			return err
		}

		// get tags from db or create them
		tags, err := createOrGetTagsFromNameList(&data.Tags, client, ctx)
		if err != nil {
			return err
		}

		// update properties
		err = createOrUpdatePropertiesFromMap(partX, &data.Properties, client, ctx)
		if err != nil {
			return err
		}

		// get box by given ID
		boxX, err := client.Box.Get(ctx, data.BoxId)
		if err != nil {
			return err
		}

		// update part with request data
		partX, err = partX.Update().
			SetName(data.Name).
			SetDescription(data.Description).
			SetAmount(amount).
			SetBox(boxX).
			ClearTags().
			AddTags(tags...).
			Save(ctx)

		if err != nil {
			return err
		}

		return c.JSON(partX)
	})

	// storage := memory.New()

	router.Put("/:partId<int>/image", func(c *fiber.Ctx) error {
		partId, _ := strconv.Atoi(c.Params("partId"))

		form, err := c.MultipartForm()
		if err != nil {
			return fiber.NewError(400, "please send a multipart form")
		}

		files := form.File["image"]
		if len(files) != 1 {
			return fiber.NewError(400, "please send exactly one file with the key 'image'")
		}

		// get partX from db
		partX, err := client.Part.Get(ctx, partId)
		if err != nil {
			return err
		}

		imageId := uuid.New()
		fp := filepath.Join(config.StoragePath, imageId.String())
		err = c.SaveFile(files[0], fp)
		if err != nil {
			return err
		}

		partX.Update().SetImageId(imageId).Exec(ctx)

		return err
	})

	router.Get("/", func(c *fiber.Ctx) error {
		// get all parts from db
		parts, err := client.Part.Query().
			WithTags().
			WithProperties().
			WithBox().
			All(ctx)
		if err != nil {
			return err
		}

		return c.JSON(parts)
	})

	router.Get("/search", func(c *fiber.Ctx) error {
		query := c.Query("q")
		filter := c.Query("filter")

		filterBy := filter
		searchResult, err := typesense_wrapper.TypesenseClient.Collection("parts").Documents().Search(&api.SearchCollectionParams{FilterBy: &filterBy, Q: query, QueryBy: "name,tags,description"})
		if err != nil {
			return err
		}

		parts := make([]*ent.Part, 0)
		for _, hit := range *searchResult.Hits {
			doc := *hit.Document
			partId, err := strconv.Atoi(doc["id"].(string))
			if err != nil {
				return err
			}

			part, err := client.Part.Query().Where(part.ID(partId)).WithTags().WithProperties().Only(ctx)
			if err != nil {
				return err
			}

			parts = append(parts, part)
		}

		return c.JSON(parts)
	})

	router.Delete("/:partId<int>", func(c *fiber.Ctx) error {
		partId, _ := strconv.Atoi(c.Params("partId"))

		// parse request body
		data := new(PartAddData)
		if err := c.BodyParser(data); err != nil {
			return err
		}

		partX, err := client.Part.UpdateOneID(partId).SetDeleted(true).Save(ctx)

		if err != nil {
			return err
		}

		return c.JSON(partX)
	})

	router.Post("/:partId<int>/deliver", func(c *fiber.Ctx) error {
		partId, _ := strconv.Atoi(c.Params("partId"))

		// get part from db
		part, err := client.Part.Get(ctx, partId)
		if err != nil {
			return err
		}

		position, err := part.QueryBox().QueryPosition().Only(ctx) //.QuerySection().QueryBox().QueryPosition().Only(ctx)
		if err != nil {
			return err
		}

		resp, err := helpers.DeliverBoxByPositionId(position.ID)
		if err != nil {
			return err
		}
		return c.JSON(resp)
	})
}
