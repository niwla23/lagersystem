package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	_ "github.com/mattn/go-sqlite3"
	"github.com/niwla23/lagersystem/manager/ent"
)

func main() {
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("sqlite3", "file:///tmp/db.sqlite?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fiber.New(fiber.Config{AppName: "Storagesystem Manager Service"})
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Storagesystem Metrics Page"}))

	app.Post("/part", func(c *fiber.Ctx) error {
		type PartAddData struct {
			Name        string   `json:"name"`
			Description string   `json:"description"`
			Tags        []string `json:"tags"`
			BoxId       int      `json:"boxId"`
		}

		data := new(PartAddData)

		if err := c.BodyParser(data); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		part, err := client.Part.Create().SetName(data.Name).SetDescription(data.Description).Save(ctx)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendString(part.Name)
	})

	app.Listen(":3001")
}

// func main() {
// 	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
// 	client, err := ent.Open("sqlite3", "file:///tmp/db.sqlite?_fk=1")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to sqlite: %v", err)
// 	}
// 	defer client.Close()
// 	ctx := context.Background()

// 	// Run the auto migration tool.
// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}

// 	pos, _ := client.Position.Create().SetPositionId(1).Save(ctx)
// 	box, _ := client.Box.Create().SetPosition(pos).Save(ctx)
// 	section, _ := client.Section.Create().SetBox(box).Save(ctx)
// 	part, _ := client.Part.Create().SetName("test part2").AddSections(section).Save(ctx)

// 	xtag, _ := client.Tag.Create().SetName("X").SetDescription("mystery tag").Save(ctx)
// 	x2tag, _ := client.Tag.Create().SetName("X2").SetDescription("x2").SetParent(xtag).Save(ctx)
// 	x3tag, _ := client.Tag.Create().SetName("X3").SetDescription("x3").SetParent(xtag).Save(ctx)

// 	part.Update().AddTags(xtag).Save(ctx)
// 	fmt.Println(part.Name, x2tag, x3tag)
// }
