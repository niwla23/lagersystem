package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "Storagesystem Operator Service"})

	app.Get("/scanBoxId", func(c *fiber.Ctx) error {
		time.Sleep(1870 * time.Millisecond)
		return c.JSON(&fiber.Map{
			"status":   "success",
			"boxId":    69,
			"duration": 1.87,
		})
	})

	app.Get("/deliver/:positionId<int>", func(c *fiber.Ctx) error {
		positionId, _ := strconv.Atoi(c.Params("positionId"))
		time.Sleep(9 * time.Second)
		return c.JSON(&fiber.Map{
			"status":   "success",
			"boxId":    positionId,
			"duration": 9.69,
		})
	})

	app.Get("/store/:positionId<int>", func(c *fiber.Ctx) error {
		positionId, _ := strconv.Atoi(c.Params("positionId"))
		time.Sleep(9 * time.Second)
		return c.JSON(&fiber.Map{
			"status":   "success",
			"boxId":    positionId,
			"duration": 9.69,
		})
	})

	app.Listen(":3000")
}
