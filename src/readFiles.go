package src

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

const asciiChars = "@%#*+=-:. "

func GetAsciiArt(file *os.File) (image.Image, []string, error) {
	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, nil, err
	}

	bounds := img.Bounds()

	art := make([]string, bounds.Max.Y)
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		line := ""
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			pixel := img.At(x, y)

			r, g, b, _ := pixel.RGBA()

			// Calculate the graysclae value (luminosity method)
			gray := uint8((.299*float64(r) + .587*float64(g) + .114*float64(b)) / 256)

			artIndex := int(gray) * (len(asciiChars) - 1) / 255

			line += string(asciiChars[artIndex])
			grayImg.Set(x, y, color.Gray{Y: gray})
		}
		art[y] = line
	}

	return grayImg, art, nil
}
