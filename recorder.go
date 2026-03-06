package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

// CaptureScreen ana ekranın görüntüsünü alır ve belirtilen isimle kaydeder.
func CaptureScreen() error {
	// Ekran sınırlarını al
	bounds := screenshot.GetDisplayBounds(0)

	// Görüntüyü yakala
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return fmt.Errorf("ekran yakalanamadı: %w", err)
	}

	// Dosya adını o anki zamanla benzersiz yapalım (Örn: screen_20260307_0230.png)
	fileName := fmt.Sprintf("screen_%s.png", time.Now().Format("20060102_150405"))
	
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("dosya oluşturulamadı: %w", err)
	}
	defer file.Close()

	// PNG olarak kaydet
	return png.Encode(file, img)
}