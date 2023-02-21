package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/niwla23/lagersystem/manager/ent/runtime"
	"github.com/niwla23/lagersystem/manager/typesense_wrapper"
	"github.com/typesense/typesense-go/typesense/api"

	"github.com/niwla23/lagersystem/manager/ent"
)

type PartAddData struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	BoxId       int      `json:"boxId"`
}

// use this function to return a json status message.
// to use it just return its error
func JsonStatusResponse(c *fiber.Ctx, code int, message string) *fiber.Error {
	schema := struct {
		Message string `json:"message"`
	}{}
	schema.Message = message

	data, err := json.Marshal(schema)

	if err != nil {
		return fiber.NewError(500, "not sure what even works if we even fail to make an error lol")
	}

	c.Response().Header.Add("content-type", "application/json")
	return fiber.NewError(code, string(data))
}

func main() {
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("sqlite3", "file:///tmp/db.sqlite?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	typesense_wrapper.InitTypesense()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	partCollection := typesense_wrapper.TypesenseClient.Collection("parts")
	_, err = partCollection.Retrieve()
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
				},
				// DefaultSortingField: "num_employees",
			}

			_, err = typesense_wrapper.TypesenseClient.Collections().Create(schema)
			if err != nil {
				panic(fmt.Sprintf("failed creating typesense collection: %v", err))
			}
		}
	}

	app := fiber.New(fiber.Config{AppName: "Storagesystem Manager Service"})
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Storagesystem Metrics Page"}))

	app.Post("/part", func(c *fiber.Ctx) error {
		data := new(PartAddData)

		if err := c.BodyParser(data); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		part, err := client.Part.Create().SetName(data.Name).SetDescription(data.Description).Save(ctx)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendString(part.Name)
	})

	// app.Put("/part/:partId<int>", func(c *fiber.Ctx) error {
	// 	partId, _ := strconv.Atoi(c.Params("partId"))
	// 	data := new(PartAddData)

	// 	if err := c.BodyParser(data); err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}

	// 	part, err := client.Part.Get(ctx, partId)

	// 	if err != nil {
	// 		return c.Status(404).SendString(err.Error())
	// 	}

	// 	part, err = part.Update().SetName(data.Name).SetDescription(data.Description).Save(ctx)

	// 	if err != nil {
	// 		return c.Status(500).SendString(err.Error())
	// 	}

	// 	responseData, err := json.Marshal(part)

	// 	if err != nil {
	// 		return JsonStatusResponse(c, fiber.StatusInternalServerError, err.Error())
	// 	}
	// 	return c.SendString(string(responseData))
	// })

	app.Listen(":3001")
}

// func main() {
// 	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
// 	client, err := ent.Open("sqlite3", "file:///tmp/db.sqlite?_fk=1")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to sqlite: %v", err)
// 	}
// 	defer client.Close()
// 	ctx := context.Background()

// 	// Run the auto migration tool.
// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}

// 	pos, _ := client.Position.Create().SetPositionId(1).Save(ctx)
// 	box, _ := client.Box.Create().SetPosition(pos).Save(ctx)
// 	section, _ := client.Section.Create().SetBox(box).Save(ctx)
// 	part, _ := client.Part.Create().SetName("test part2").AddSections(section).Save(ctx)

// 	xtag, _ := client.Tag.Create().SetName("X").SetDescription("mystery tag").Save(ctx)
// 	x2tag, _ := client.Tag.Create().SetName("X2").SetDescription("x2").SetParent(xtag).Save(ctx)
// 	x3tag, _ := client.Tag.Create().SetName("X3").SetDescription("x3").SetParent(xtag).Save(ctx)

// 	part.Update().AddTags(xtag).Save(ctx)
// 	fmt.Println(part.Name, x2tag, x3tag)
// }
