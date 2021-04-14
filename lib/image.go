package lib

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

type ImageBuilder struct {
}

func (builder ImageBuilder) Make(filename string) image.Image {
	image, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	decoded, err := jpeg.Decode(image)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}

	defer image.Close()

	return decoded
}

func (builder ImageBuilder) Insert(source image.Image, target image.Image, x int, y int) image.Image {
	offset := image.Pt(x, y)
	bounds := source.Bounds()
	output := image.NewRGBA(bounds)
	draw.Draw(output, bounds, source, image.ZP, draw.Src)
	draw.Draw(output, target.Bounds().Add(offset), target, image.ZP, draw.Over)

	return output
}

func (builder ImageBuilder) Save(image image.Image, filename string) {
	output, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}

	jpeg.Encode(output, image, &jpeg.Options{jpeg.DefaultQuality})

	defer output.Close()
}
