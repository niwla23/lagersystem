package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
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
			SetID(data.BoxId).
			Save(ctx)

		if err != nil {
			return err
		}

		return c.JSON(boxX)
	})

	router.Get("/", func(c *fiber.Ctx) error {
		// get all boxes from db
		boxes, err := client.Box.Query().WithPosition().WithParts().All(ctx)
		if err != nil {
			return err
		}

		return c.JSON(boxes)
	})

	router.Get("/get-free-box", func(c *fiber.Ctx) error {
		box, err := client.Box.Query(). // query a box
						Where(box.HasPositionWith(position.HasWarehouseWith(warehouse.ID(1)))). // that is currently stored in warehouse 1
						Where(box.Not(box.HasParts())).                                         // that has no parts in it
						WithPosition().                                                         // include the position of the box
						First(ctx)                                                              // get the first result
		if err != nil {
			return err
		}
		return c.JSON(box)
	})

	router.Get("/:boxId", func(c *fiber.Ctx) error {
		boxId, err := uuid.Parse(c.Params("boxId"))
		if err != nil {
			return err
		}
		boxX, err := client.Box.Query().Where(box.ID(boxId)).WithParts().WithPosition().Only(ctx)
		if err != nil {
			return err
		}
		return c.JSON(boxX)
	})

	// router.Get("/get-by-scanner", func(c *fiber.Ctx) error {
	// 	boxX, err := helpers.ScanIoPos()
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return c.JSON(boxX)
	// })

	router.Post("/:boxId/deliver", func(c *fiber.Ctx) error {
		boxId, err := uuid.Parse(c.Params("boxId"))
		if err != nil {
			return err
		}
		// ioPos := c.Params("ioPos")

		// ioState, err := helpers.GetIOState()
		// if err != nil {
		// 	return err
		// }

		// isWantedIoOk := helpers.IsIoSlotFree(ioState, ioPos)

		// if ioPos == "0" && !isWantedIoOk {
		// 	ioPos, err = helpers.FindIoSlot()
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		// get box from db
		box, err := client.Box.Query().Where(box.ID(boxId)).Only(ctx)
		if err != nil {
			return err
		}

		// get position of box from db
		position, err := box.QueryPosition().Only(ctx)
		if err != nil {
			return err
		}

		resp, err := helpers.PickupBox(position.ID)
		if err != nil {
			return err
		}

		// update position of box
		_, err = client.Box.UpdateOne(box).ClearPosition().Save(ctx)
		if err != nil {
			return err
		}

		return c.JSON(resp)
	})

	router.Get("/getFromScanner", func(c *fiber.Ctx) error {
		startTime := time.Now()

		ioState, err := helpers.GetIOState()
		if err != nil {
			return err
		}

		if helpers.IsIoSlotFree(ioState, "1") {
			return errors.New("no box in scanner")
		}

		// scan box and find it in db
		boxX, _, err := helpers.ScanIoPos("1")
		if err != nil {
			return err
		}

		return c.JSON(&fiber.Map{
			"status":   "success",
			"boxId":    boxX.ID,
			"duration": time.Since(startTime).Seconds(),
		})
	})
}
