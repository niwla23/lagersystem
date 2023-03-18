package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	ent_gen "github.com/niwla23/lagersystem/manager/ent/generated"
)

type WarehouseAddData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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
		// fmt.Println(warehouses)

		return c.JSON(warehouses)
	})
}
