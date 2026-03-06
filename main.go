package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Watchbro İzleme Servisi Başlatıldı...")

	// 10 saniyede bir tetiklenecek bir zamanlayıcı oluştur
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Programı durdurana kadar döngüde kalsın
	for range ticker.C {
		fmt.Println("Ekran görüntüsü alınıyor...")
		
		err := CaptureScreen()
		if err != nil {
			fmt.Printf("Hata oluştu: %v\n", err)
		} else {
			fmt.Println("Görüntü başarıyla kaydedildi.")
		}
	}
	
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/autorun/on", func(c *fiber.Ctx) error {
		AutoRunEnable()
		return c.SendString("Autorun is on")
	})

	app.Get("/autorun/off", func(c *fiber.Ctx) error {
		AutoRunDisable()
		return c.SendString("Autorun is off")
	})

	app.Get("/wifi/on", func(c *fiber.Ctx) error {
		wifiOnOff(true)
		return c.SendString("Wifi is on")
	})

	app.Get("/wifi/off", func(c *fiber.Ctx) error {
		wifiOnOff(false)
		return c.SendString("Wifi is off")
	})

	app.Get("/screenshot", func(c *fiber.Ctx) error {
		takeScreenshot()
		return c.SendString("Screenshot taken")
	})

	fmt.Println("Server is running on port 3000")
	app.Listen(":3000")
}
