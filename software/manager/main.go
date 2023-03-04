package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/niwla23/lagersystem/manager/ent/runtime"
	"github.com/niwla23/lagersystem/manager/typesense_wrapper"
	"github.com/typesense/typesense-go/typesense/api"

	"github.com/niwla23/lagersystem/manager/ent"
)

// use this function to return a json status message.
// to use it just return its error
// DEPRECATED: just yeet the error directly, gofiber will take care of it
func JsonStatusResponse(c *fiber.Ctx, code int, message string) error {
	schema := struct {
		Message   string    `json:"message"`
		Code      int       `json:"code"`
		Timestamp time.Time `json:"timestamp"`
	}{}
	schema.Message = message
	schema.Code = code
	schema.Timestamp = time.Now().UTC()

	data, err := json.Marshal(schema)

	if err != nil {
		return fiber.NewError(500, "not sure what even works if we even fail to make an error lol")
	}

	c.Response().Header.Add("content-type", "application/json")
	return c.Status(code).SendString(string(data))
	// return fiber.NewError(code, string(data))
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
		} else {
			panic(fmt.Sprintf("failed getting typesense collection: %v", err))
		}
	}

	app := fiber.New(fiber.Config{
		AppName: "Storagesystem Manager Service",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			fmt.Println(err.Error())

			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// 404 if its a entgo not found error
			if _, ok := err.(*ent.NotFoundError); ok {
				code = fiber.StatusNotFound
			}

			// 406 if its a entgo constraint error
			if _, ok := err.(*ent.ConstraintError); ok {
				code = fiber.StatusNotAcceptable
			}

			// 406 if its a entgo validation error
			if _, ok := err.(*ent.ValidationError); ok {
				code = fiber.StatusNotAcceptable
			}

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return JsonStatusResponse(c, code, err.Error())
		},
	})

	app.Use(recover.New(recover.Config{EnableStackTrace: true}))

	partHandlers := app.Group("/parts")
	registerPartRoutes(partHandlers, client, ctx)

	positionHandlers := app.Group("/positions")
	registerPositionRoutes(positionHandlers, client, ctx)

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
