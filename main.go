package main

import (
	"fmt"
	"os"

	"images/src"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . lion.jpg lion")
		return
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	defer file.Close()

	img, art, err := src.GetAsciiArt(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputFile,err := src.DownloadFiles(img,art,fileName)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Printf("GrayScale image output as %v.png\n",outputFile)
	fmt.Printf("File exported as %v.txt\n", outputFile)
}
