package main

import (
	"fmt"
	"image"
	"image/draw"

	_ "image/jpeg"
	"image/png"
	"os"
)

func main() {
	// Define file paths for background and foreground images
	bgImagePath := "346_digital_marketing_social_media_post_template.png"
	fgImagePath := "3footer.png"

	bgFile, err := os.Open(bgImagePath)
	if err != nil {
		fmt.Println("Error opening background image:", err)
		return
	}
	defer bgFile.Close()

	bgImg, err := png.Decode(bgFile)
	if err != nil {
		fmt.Println("Error decoding background image:", err)
		return
	}

	fgFile, err := os.Open(fgImagePath)
	if err != nil {
		fmt.Println("Error opening foreground image:", err)
		return
	}
	defer fgFile.Close()

	fgImg, err := png.Decode(fgFile)
	if err != nil {
		fmt.Println("Error decoding foreground image:", err)
		return
	}

	// Get the bounds of the background and foreground images
	bgBounds := bgImg.Bounds()
	fgBounds := fgImg.Bounds()

	// Define the target image size (assuming foreground fits within background)
	targetWidth := bgBounds.Max.X
	targetHeight := bgBounds.Max.Y

	// Create a new RGBA image for the final result
	resultImg := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))

	// Draw the background image to the result
	draw.Draw(resultImg, resultImg.Bounds(), bgImg, bgBounds.Min, draw.Src)

	// Calculate the offset for the foreground image (place it at the bottom)
	fgOffsetX := bgBounds.Min.X + (bgBounds.Max.X-fgBounds.Max.X)/2 // Center horizontally
	fgOffsetY := bgBounds.Max.Y - fgBounds.Max.Y                    // Place at bottom

	// Draw the foreground image to the result with offset
	draw.Draw(resultImg, fgBounds.Add(image.Point{fgOffsetX, fgOffsetY}), fgImg, fgBounds.Min, draw.Over)

	// Save the final image
	outputFile, err := os.Create("combined3.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, resultImg)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}

	fmt.Println("Successfully combined images!")
}
