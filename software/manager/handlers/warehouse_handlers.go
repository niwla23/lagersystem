package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	ent_gen "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/helpers"
)

type WarehouseAddData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ClearIoPos(ioPosId string, client *ent_gen.Client, ctx context.Context, boxX *ent_gen.Box) (*ent_gen.Box, error) {
	ioState, err := helpers.GetIOState()
	if err != nil {
		return nil, err
	}

	// scanner must be free if we are not storing from scanner
	if !helpers.IsIoSlotFree(ioState, "1") && ioPosId != "1" {
		return nil, errors.New("scanner not free")
	}

	// source io pos must be occupied
	if helpers.IsIoSlotFree(ioState, ioPosId) {
		return nil, errors.New("source io pos not occupied")
	}

	// scan box and find it in db

	if boxX == nil {
		boxX, _, err = helpers.ScanIoPos(ioPosId)
		if err != nil {
			return boxX, err
		}
	}

	// find a position for the box
	positionX, err := helpers.FindPosition(boxX)
	if err != nil {
		return boxX, err
	}

	// at this point we definetely have a free position for the box and can tell the operator service to store it there
	helpers.StoreBox(positionX.ID, "1")
	client.Box.UpdateOne(boxX).SetPosition(positionX).Save(ctx)

	return boxX, nil
}

func RegisterWarehouseRoutes(router fiber.Router, client *ent_gen.Client, ctx context.Context) {
	router.Post("/", func(c *fiber.Ctx) error {
		data := new(WarehouseAddData)

		if err := c.BodyParser(data); err != nil {
			return err
		}

		warehouseX, err := client.Warehouse.Create().
			SetName(data.Name).
			SetDescription(data.Description).
			Save(ctx)

		if err != nil {
			return err
		}

		return c.JSON(warehouseX)
	})

	router.Get("/", func(c *fiber.Ctx) error {
		// get all warehouses from db
		warehouses, err := client.Warehouse.Query().All(ctx)
		if err != nil {
			return err
		}

		return c.JSON(warehouses)
	})

	router.Post("/:warehouseId<int>/clearIO", func(c *fiber.Ctx) error {
		startTime := time.Now()

		ioState, err := helpers.GetIOState()
		if err != nil {
			return err
		}

		// check for each io pos if it is free. If not, clear it
		clearedIds := make([]string, 0)
		for i := 1; i <= 3; i++ {
			if !helpers.IsIoSlotFree(ioState, fmt.Sprint(i)) {
				_, err := ClearIoPos(fmt.Sprint(i), client, ctx, nil)
				if err != nil {
					return err
				}
				clearedIds = append(clearedIds, fmt.Sprint(i))
			}
		}

		return c.JSON(&fiber.Map{
			"status":   "success",
			"cleared":  clearedIds,
			"duration": time.Since(startTime).Seconds(),
		})
	})

	router.Post("/:warehouseId<int>/clearScanner", func(c *fiber.Ctx) error {
		startTime := time.Now()

		ioState, err := helpers.GetIOState()
		if err != nil {
			return err
		}

		if helpers.IsIoSlotFree(ioState, "1") {
			return errors.New("source is free")
		}

		boxX, err := ClearIoPos("1", client, ctx, nil)
		if err != nil {
			return err
		}

		return c.JSON(&fiber.Map{
			"status":   "success",
			"boxId":    boxX.ID,
			"duration": time.Since(startTime).Seconds(),
		})
	})

	router.Get("/1/getPositions", func(c *fiber.Ctx) error {
		positions, err := helpers.GetPositions(client)

		if err != nil {
			return err
		}

		return c.JSON(positions)
	})

	router.Get("/1/getIOState", func(c *fiber.Ctx) error {
		ioState, err := helpers.GetIOState()

		if err != nil {
			return err
		}

		return c.JSON(ioState)
	})
}
