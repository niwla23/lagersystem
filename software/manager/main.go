package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/mattn/go-sqlite3"

	"github.com/niwla23/lagersystem/manager/config"
	"github.com/niwla23/lagersystem/manager/database"
	_ "github.com/niwla23/lagersystem/manager/ent/generated/runtime"
	"github.com/niwla23/lagersystem/manager/typesense_sync"
	"github.com/niwla23/lagersystem/manager/typesense_wrapper"

	ent "github.com/niwla23/lagersystem/manager/ent/generated"

	"github.com/niwla23/lagersystem/manager/handlers"
)

func main() {
	config.LoadConfigFromEnvironment()

	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	// client, err := ent.Open("sqlite3", config.DBUri)
	// if err != nil {
	// 	log.Fatalf("failed opening connection to sqlite: %v", err)
	// }
	// defer client.Close()
	database.InitDB()
	ctx := context.Background()

	typesense_wrapper.InitTypesense()

	// // Run the auto migration tool.
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	if !fiber.IsChild() {
		go typesense_sync.SyncBackgroundTask()
	}

	app := fiber.New(fiber.Config{
		AppName:   "Storagesystem Manager Service",
		BodyLimit: 100 * 1024 * 1024, // size in MB,
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

			return c.Status(code).JSON(&fiber.Map{
				"message":   err.Error(),
				"code":      code,
				"timestamp": time.Now().UTC(),
			})
		},
	})

	app.Use(requestid.New())
	app.Use(favicon.New())
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(etag.New())

	app.Static("/static", config.StoragePath)

	partHandlers := app.Group("/parts")
	handlers.RegisterPartRoutes(partHandlers, database.Client, ctx)

	positionHandlers := app.Group("/positions")
	handlers.RegisterPositionRoutes(positionHandlers, database.Client, ctx)

	boxHandlers := app.Group("/boxes")
	handlers.RegisterBoxRoutes(boxHandlers, database.Client, ctx)

	tagHandlers := app.Group("/tags")
	handlers.RegisterTagRoutes(tagHandlers, database.Client, ctx)

	warehouseHandlers := app.Group("/warehouses")
	handlers.RegisterWarehouseRoutes(warehouseHandlers, database.Client, ctx)

	app.Listen(":3001")
}
