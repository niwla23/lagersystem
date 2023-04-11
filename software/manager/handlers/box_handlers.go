package handlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
	"github.com/niwla23/lagersystem/manager/ent/generated/section"
	"github.com/niwla23/lagersystem/manager/ent/generated/warehouse"
	"github.com/niwla23/lagersystem/manager/helpers"
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

	router.Get("/get-free-box", func(c *fiber.Ctx) error {
		box, err := client.Box.Query(). // query a box
						Where(box.HasPositionWith(position.HasWarehouseWith(warehouse.ID(1)))). // that is currently stored in warehouse 1
						Where(box.Not(box.HasSectionsWith(section.HasParts()))).                // that has no parts in it
						WithPosition().                                                         // include the position of the box
						Only(ctx)                                                               // get the first result
		if err != nil {
			return err
		}
		return c.JSON(box)
	})

	router.Get("/get-by-scanner", func(c *fiber.Ctx) error {
		boxX, err := helpers.GetBoxFromScanner()
		if err != nil {
			return err
		}
		return c.JSON(boxX)
	})

	router.Post("/:boxId<int>/deliver", func(c *fiber.Ctx) error {
		boxId, _ := strconv.Atoi(c.Params("boxId"))

		// get box from db
		box, err := client.Box.Get(ctx, boxId)
		if err != nil {
			return err
		}
		position, err := box.QueryPosition().Only(ctx)
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
