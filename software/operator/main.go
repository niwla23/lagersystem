package main

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gofiber/fiber/v2"
	"github.com/kdar/goquirc"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "Storagesystem Operator Service"})

	app.Get("/scanBoxId", func(c *fiber.Ctx) error {
		// file, err := os.Open("qrcode.jpg")
		// if err != nil {
		// 	return fiber.NewError(fiber.StatusInternalServerError, "cant read image file")
		// }

		// img, format, err := image.Decode(file)
		// if err != nil {
		// 	fmt.Println(err, format)
		// 	return fiber.NewError(fiber.StatusInternalServerError, "cant decode image file")
		// }

		// bmp, err := gozxing.NewBinaryBitmapFromImage(img)
		// if err != nil {
		// 	return fiber.NewError(fiber.StatusInternalServerError, "cant convert image file")
		// }

		// // decode image
		// qrReader := qrcode.NewQRCodeReader()
		// result, err := qrReader.Decode(bmp, nil)
		// if err != nil {
		// 	return fiber.NewError(fiber.StatusInternalServerError, "cant read qr code")
		// }

		// boxId := result.GetText()

		imgdata, err := ioutil.ReadFile("qrcode.jpg")
		if err != nil {
			log.Fatal(":", err)
		}

		// Decode image
		m, _, err := image.Decode(bytes.NewReader(imgdata))
		if err != nil {
			log.Fatal(":", err)
		}

		d := goquirc.New()
		defer d.Destroy()
		datas, err := d.Decode(m)
		if err != nil {
			log.Fatal(":", err)
		}

		for _, data := range datas {
			fmt.Printf("%s\n", data.Payload[:data.PayloadLen])
		}

		boxId := "b1"

		return c.JSON(&fiber.Map{
			"status":   "success",
			"boxId":    boxId,
			"duration": 1.87,
		})
	})

	app.Get("/deliver/:positionId<int>", func(c *fiber.Ctx) error {
		positionId, _ := strconv.Atoi(c.Params("positionId"))
		time.Sleep(9 * time.Second)
		return c.JSON(&fiber.Map{
			"status":     "success",
			"positionId": positionId,
			"duration":   9.69,
		})
	})

	app.Get("/store/:positionId<int>", func(c *fiber.Ctx) error {
		positionId, _ := strconv.Atoi(c.Params("positionId"))
		time.Sleep(9 * time.Second)
		return c.JSON(&fiber.Map{
			"status":     "success",
			"positionId": positionId,
			"duration":   9.69,
		})
	})

	app.Listen(":3000")
}
