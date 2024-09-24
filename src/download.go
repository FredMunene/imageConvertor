package src

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
)

func DownloadFiles(img image.Image, art []string,fileName string ) (string,error) {

	outputFileName := strings.Split(fileName,".")[0]
	if outputFileName == ""{
		return "", fmt.Errorf(" Error: empty file name")
	}
	outputFileTxt, err := os.Create(outputFileName + ".txt")
	if err != nil {
		fmt.Println("Error creating txt file", err)
		return "", err
	}
	outputFilePng, err := os.Create(outputFileName + ".png")
	if err != nil {
		fmt.Println("Error creating image file", err)
		return "", err
	}

	err = png.Encode(outputFilePng, img)
	if err != nil {
		fmt.Println("Error encoding img", err)
		return "", err

	}

	for _, line := range art {
		_, err := outputFileTxt.WriteString((line + "\n"))
		if err != nil {
			
			return "", err
		}
	}

	return outputFileName,nil
}
