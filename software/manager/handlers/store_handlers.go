package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	// "github.com/niwla23/lagersystem/manager/ent/box"
	ent_gen "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/helpers"
	// "github.com/niwla23/lagersystem/manager/ent/position"
	// "github.com/niwla23/lagersystem/manager/ent/warehouse"
)

func RegisterStoreRoutes(router fiber.Router, client *ent_gen.Client, ctx context.Context) {
	router.Post("/by-scanner", func(c *fiber.Ctx) error {
		boxX, err := helpers.GetBoxFromScanner()
		if err != nil {
			return err
		}

		positionX, err := helpers.FindPosition(boxX)
		if err != nil {
			return err
		}

		// at this point we definetely have a free position for the box and can tell the operator service to store it there
		fmt.Println("in another universe we would ask operator service to store box at position: " + fmt.Sprint(positionX.ID))

		return c.JSON(&fiber.Map{
			"status":     "success",
			"boxId":      boxX.BoxId,
			"positionId": positionX.ID,
			"duration":   1.87,
		})
	})
}
