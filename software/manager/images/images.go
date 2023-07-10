package images

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"

	"github.com/kolesa-team/go-webp/webp"
	"golang.org/x/image/draw"
)

// Helper function to calculate the scaled dimensions while maintaining aspect ratio
func calculateScaledDimensions(originalWidth, originalHeight, maxWidth, maxHeight int) (int, int) {
	aspectRatio := float64(originalWidth) / float64(originalHeight)

	if originalWidth <= maxWidth && originalHeight <= maxHeight {
		// No scaling needed, return original dimensions
		return originalWidth, originalHeight
	}

	width := maxWidth
	height := int(float64(width) / aspectRatio)

	if height > maxHeight {
		height = maxHeight
		width = int(float64(height) * aspectRatio)
	}

	return width, height
}

// Helper function to scale the image to the specified dimensions
func scaleImage(img image.Image, width, height int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.BiLinear.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)
	return dst
}

// returns the new filename
func SaveImage(header *multipart.FileHeader, fp string) (string, error) {
	file, err := header.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// Calculate the scaled dimensions while maintaining aspect ratio
	maxWidth := 2000
	maxHeight := 2000
	width, height := calculateScaledDimensions(img.Bounds().Dx(), img.Bounds().Dy(), maxWidth, maxHeight)

	// Scale the image if necessary
	if width != img.Bounds().Dx() && height != img.Bounds().Dy() {
		img = scaleImage(img, width, height)
	}

	// Rename file to .webp
	outFile, err := os.Create(fp)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Encode the image as WebP and save it to the temporary file
	err = webp.Encode(outFile, img, nil)
	if err != nil {
		return "", err
	}

	newFilename := fmt.Sprintf("%s.webp", header.Filename)

	return newFilename, nil
}

func IsImage(header multipart.FileHeader) bool {
	file, err := header.Open()
	if err != nil {
		return false
	}
	defer file.Close()

	// Decode the image
	_, _, err = image.Decode(file)
	return err == nil
}
