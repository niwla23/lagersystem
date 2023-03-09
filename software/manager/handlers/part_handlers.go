package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/part"
	"github.com/niwla23/lagersystem/manager/ent/generated/property"
	"github.com/niwla23/lagersystem/manager/ent/generated/tag"
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

		part, err := client.Part.Create().
			SetName(data.Name).
			SetDescription(data.Description).
			SetAmount(amount).
			AddTags(tags...).
			Save(ctx)

		if err != nil {
			return err
		}

		// add properties
		err = createOrUpdatePropertiesFromMap(part, &data.Properties, client, ctx)
		if err != nil {
			return err
		}

		// encode part to json
		responseData, err := json.Marshal(part)
		if err != nil {
			return err
		}
		return c.SendString(string(responseData))
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

		// get part from db
		part, err := client.Part.Query().Where(part.ID(partId)).WithTags().WithProperties().Only(ctx)
		if err != nil {
			return err
		}

		// get tags from db or create them
		tags, err := createOrGetTagsFromNameList(&data.Tags, client, ctx)
		if err != nil {
			return err
		}

		// update properties
		err = createOrUpdatePropertiesFromMap(part, &data.Properties, client, ctx)
		if err != nil {
			return err
		}

		// update part with request data
		part, err = part.Update().
			SetName(data.Name).
			SetDescription(data.Description).
			SetAmount(amount).
			ClearTags().
			AddTags(tags...).
			Save(ctx)

		if err != nil {
			return err
		}

		// encode part to json
		responseData, err := json.Marshal(part)
		if err != nil {
			return err
		}
		return c.SendString(string(responseData))
	})

	router.Get("/", func(c *fiber.Ctx) error {
		// get all parts from db
		parts, err := client.Part.Query().WithTags().WithProperties().All(ctx)
		if err != nil {
			return err
		}

		// encode parts to json
		responseData, err := json.Marshal(parts)
		if err != nil {
			return err
		}
		return c.SendString(string(responseData))
	})

	router.Post("/:partId<int>/deliver", func(c *fiber.Ctx) error {
		partId, _ := strconv.Atoi(c.Params("partId"))

		// get part from db
		part, err := client.Part.Get(ctx, partId)
		if err != nil {
			return err
		}
		position, err := part.QuerySection().QueryBox().QueryPosition().Only(ctx)
		if err != nil {
			return err
		}

		return c.SendString("in another universe we would ask operator service to deliver box at position: " + fmt.Sprint(position.ID))
	})

	// THIS IS PROBABLY COMPLETLY USELESS AS YOU WANT TO STORE THE BOX PLACED ON THE READER AND NOT THE PART
	// router.Post("/:partId<int>/store", func(c *fiber.Ctx) error {
	// 	partId, _ := strconv.Atoi(c.Params("partId"))

	// 	// get part from db
	// 	part, err := client.Part.Get(ctx, partId)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	positionX, err := part.QuerySection().QueryBox().QueryPosition().Only(ctx)

	// 	target := &ent.NotFoundError{}
	// 	if errors.As(err, &target) {
	// 		// find free position
	// 		positionX, err = client.Position.Query().
	// 			Where(position.HasWarehouseWith(warehouse.ID(1))).
	// 			Where(position.Not(position.HasStoredBox())).
	// 			Only(ctx)

	// 		if err != nil {
	// 			return err
	// 		}
	// 	} else if err != nil {
	// 		return err
	// 	}

	// 	return c.SendString("in another universe we would ask operator service to store part at loaction: " + fmt.Sprint(positionX.ID))
	// })
}
