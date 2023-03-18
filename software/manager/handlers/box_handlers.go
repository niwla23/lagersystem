package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"
)

type BoxAddData struct {
	BoxId uuid.UUID `json:"boxId"`
}

func RegisterBoxRoutes(router fiber.Router, client *ent.Client, ctx context.Context) {
	// create a new box
	router.Post("/", func(c *fiber.Ctx) error {
		data := new(BoxAddData)

		if err := c.BodyParser(data); err != nil {
			return err
		}

		boxX, err := client.Box.Create().
			SetBoxId(data.BoxId).
			Save(ctx)

		if err != nil {
			return err
		}

		return c.JSON(boxX)
	})

	router.Get("/", func(c *fiber.Ctx) error {
		// get all boxes from db
		boxes, err := client.Box.Query().WithPosition().All(ctx)
		if err != nil {
			return err
		}

		return c.JSON(boxes)
	})

}
