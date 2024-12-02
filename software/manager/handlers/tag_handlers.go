package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"

	// "github.com/niwla23/lagersystem/manager/ent/box"
	ent_gen "github.com/niwla23/lagersystem/manager/ent/generated"
	// "github.com/niwla23/lagersystem/manager/ent/position"
	// "github.com/niwla23/lagersystem/manager/ent/warehouse"
)

func RegisterTagRoutes(router fiber.Router, client *ent_gen.Client, ctx context.Context) {
	router.Get("/", func(c *fiber.Ctx) error {
		// get all tags from db
		tags, err := client.Tag.Query().All(ctx)
		if err != nil {
			return err
		}
		return c.JSON(tags)
	})
}
