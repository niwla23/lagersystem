package handlers

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/niwla23/lagersystem/manager/ent"
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

		part, err := client.Box.Create().
			SetBoxId(data.BoxId).
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
}
