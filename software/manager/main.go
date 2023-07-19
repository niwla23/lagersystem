package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/niwla23/lagersystem/manager/config"
	"github.com/niwla23/lagersystem/manager/database"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"

	_ "github.com/niwla23/lagersystem/manager/ent/generated/runtime"
	"github.com/niwla23/lagersystem/manager/handlers"
	"github.com/niwla23/lagersystem/manager/typesense_sync"
	"github.com/niwla23/lagersystem/manager/typesense_wrapper"
)

//go:embed fake_frontend/*
var frontendFs embed.FS

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

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(frontendFs),
		Browse:     false,
		PathPrefix: "fake_frontend",
	}))

	app.Use(func(c *fiber.Ctx) error {
		// Check if the route is a backend route (starting with /api)
		c.Status(200)
		if strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}

		content, err := frontendFs.Open("fake_frontend/index.html")
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Frontend not found")
		}

		c.Type("text/html")
		c.Set("content-type", "text/html; charset=utf-8")
		c.Set("Content-Disposition", "inline")
		c.Status(200)
		return c.SendStream(content)
	})

	app.Static("/api/static", config.StoragePath)

	partHandlers := app.Group("/api/parts")
	handlers.RegisterPartRoutes(partHandlers, database.Client, ctx)

	positionHandlers := app.Group("/api/positions")
	handlers.RegisterPositionRoutes(positionHandlers, database.Client, ctx)

	boxHandlers := app.Group("/api/boxes")
	handlers.RegisterBoxRoutes(boxHandlers, database.Client, ctx)

	tagHandlers := app.Group("/api/tags")
	handlers.RegisterTagRoutes(tagHandlers, database.Client, ctx)

	warehouseHandlers := app.Group("/api/warehouses")
	handlers.RegisterWarehouseRoutes(warehouseHandlers, database.Client, ctx)

	app.Listen(":3001")
}
