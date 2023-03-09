package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	// "github.com/niwla23/lagersystem/manager/ent/box"
	ent_gen "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
	"github.com/niwla23/lagersystem/manager/ent/generated/warehouse"
	// "github.com/niwla23/lagersystem/manager/ent/position"
	// "github.com/niwla23/lagersystem/manager/ent/warehouse"
)

type scanBoxIdResponse struct {
	Status   string    `json:"status"`
	BoxId    uuid.UUID `json:"boxId"`
	Duration float64   `json:"duration"`
}

func RegisterStoreRoutes(router fiber.Router, client *ent_gen.Client, ctx context.Context) {
	router.Post("/by-scanner", func(c *fiber.Ctx) error {
		// get boxId from operator service
		resp, err := http.Get("http://localhost:3000/scanBoxId")
		if err != nil {
			return err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		data := new(scanBoxIdResponse)
		json.Unmarshal(body, data)

		// get boxX from db
		boxX, err := client.Box.Query().Where(box.BoxId(data.BoxId)).Only(ctx)
		if err != nil {
			return err
		}

		// try finding position of box in the database
		positionX, err := boxX.QueryPosition().Only(ctx)

		target := &ent_gen.NotFoundError{}
		if errors.As(err, &target) {
			// we have no position stored, find one
			positionX, err = client.Position.Query().
				Where(position.HasWarehouseWith(warehouse.ID(1))).
				Where(position.Not(position.HasStoredBox())).
				Only(ctx)

			if err != nil {
				return fiber.NewError(http.StatusNotFound, "no free position found for box: "+fmt.Sprint(boxX.BoxId)+", "+err.Error())
			}
		} else if err != nil {
			// shit has gone terribly wrong
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

// router.Post("/:partId<int>/store", func(c *fiber.Ctx) error {
// partId, _ := strconv.Atoi(c.Params("partId"))

// // get part from db
// part, err := client.Part.Get(ctx, partId)
// if err != nil {
// 	return err
// }
// positionX, err := part.QuerySection().QueryBox().QueryPosition().Only(ctx)

// target := &ent.NotFoundError{}
// if errors.As(err, &target) {
// 	// find free position
// 	positionX, err = client.Position.Query().
// 		Where(position.HasWarehouseWith(warehouse.ID(1))).
// 		Where(position.Not(position.HasStoredBox())).
// 		Only(ctx)

// 	if err != nil {
// 		return err
// 	}
// } else if err != nil {
// 	return err
// }

// return c.SendString("in another universe we would ask operator service to store part at loaction: " + fmt.Sprint(positionX.ID))
// })
