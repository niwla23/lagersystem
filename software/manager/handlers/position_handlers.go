package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	ent_gen "github.com/niwla23/lagersystem/manager/ent/generated"
)

type PositionAddData struct {
	Amount      int `json:"amount"`
	WarehouseId int `json:"warehouseId"`
}

func RegisterPositionRoutes(router fiber.Router, client *ent_gen.Client, ctx context.Context) {
	router.Post("/createAmount", func(c *fiber.Ctx) error {
		data := new(PositionAddData)

		if err := c.BodyParser(data); err != nil {
			return err
		}

		warehouse, err := client.Warehouse.Get(ctx, data.WarehouseId)

		if err != nil {
			return err
		}

		for i := 0; i < data.Amount; i++ {
			_, err := client.Position.Create().
				SetWarehouse(warehouse).
				Save(ctx)

			if err != nil {
				return err
			}

		}

		return c.JSON(&fiber.Map{"success": true, "message": "Created " + fmt.Sprint(data.Amount) + " positions"})
	})

	router.Get("/", func(c *fiber.Ctx) error {
		// get all positions from db
		positions, err := client.Position.Query().WithStoredBox().WithWarehouse().All(ctx)
		if err != nil {
			return err
		}

		return c.JSON(positions)
	})
}
