package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

// get how images are read
// difference between bright colors and dull
// brightness, shadows, etc
// reassign pixels with letters
//
const asciiChars = "@%#*+=-:. " // Characters from darkest to lightest
func main() {
	// Decode the JPEG data.
	fmt.Println("here")
	// create a reader : readonly

	file, err := os.Open("lion.jpg")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	defer file.Close()

	// data := []byte{}
	// data, err := io.ReadAll(reader)
	// if err != nil {
	// 	fmt.Println("Error reading data", err)
	// 	return
	// }
	// fmt.Println(data)

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Error decoding img", err)
		return
	}

	// pixels as an image.Image object
	// each pixel is a cordinate(x,y) within the image's bounds.
	// pixel values can be retrieved as color.Color types, using RGBA

	// loop through the image pixels
	bounds := img.Bounds()

	// a slice to hold the Ascii art
	art := make([]string, bounds.Max.Y)
	grayImg := image.NewGray(bounds)
	// y = 0, y < 420,
	for y := bounds.Min.Y; y < bounds.Max.Y; y++{
		line := ""
		for x := bounds.Min.X; x < bounds.Max.X; x++{
			// Get color of pixel
			pixel := img.At(x,y)

			// convert color to RGBA
			r,g,b, _ := pixel.RGBA()

			// Calculate the graysclae value (luminosity method)
			gray := uint8((.299*float64(r) + .587*float64(g) + .114*float64(b)) / 256)
			// map grayscale value to art character
			artIndex := int(gray) * (len(asciiChars) - 1) / 255
			// set the grayscale value to the new image
			line += string(asciiChars[artIndex])
			grayImg.Set(x,y,color.Gray{Y: gray})
		}
		art[y] = line
	}

	// Save the grayscale image to a new file
	fmt.Println("here")
	outputFile, err := os.Create("lion_bw.txt")
	if err != nil {
		fmt.Println("Error decoding img", err)
		return
	}


	// defer outputFile.Close()

	// err = jpeg.Encode(outputFile,grayImg,nil)
	// if err != nil {
	// 	fmt.Println("Error decoding img", err)
	// 	return
	// }

	fmt.Println("here")

	for _, line := range art{
		_, err := outputFile.WriteString((line + "\n"))
		if err != nil{
			fmt.Println("Error writing to file:", err)
			return
		}
		
	}

	fmt.Printf("Image format : %T\n", img)
	fmt.Printf("Image bounds : %v\n", img.Bounds()) // dimensions of image
}
