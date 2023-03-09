package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	ent_gen "github.com/niwla23/lagersystem/manager/ent/generated"
)

type PositionAddData struct {
	PositionId  int `json:"positionId"`
	WarehouseId int `json:"warehouseId"`
}

func RegisterPositionRoutes(router fiber.Router, client *ent_gen.Client, ctx context.Context) {
	router.Post("/", func(c *fiber.Ctx) error {
		data := new(PositionAddData)

		if err := c.BodyParser(data); err != nil {
			return err
		}

		warehouse, err := client.Warehouse.Get(ctx, data.WarehouseId)

		if err != nil {
			return err
		}

		positionX, err := client.Position.Create().
			SetPositionId(data.PositionId).
			SetWarehouse(warehouse).
			Save(ctx)

		if err != nil {
			return err
		}

		return c.SendString(fmt.Sprint(positionX.PositionId))
	})
}
