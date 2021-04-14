package lib

import (
	"errors"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

type ImageBuilder struct {
}

func (builder ImageBuilder) IsImageURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (builder ImageBuilder) DownloadImageURL(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (builder ImageBuilder) Decode(filename string, img io.Reader) (image.Image, error) {
	if builder.IsImageURL(filename) {
		return jpeg.Decode(img)
	}

	mime, err := mimetype.DetectFile(filename)

	if err != nil {
		return nil, err
	}

	switch mime := mime.String(); mime {
	case "image/png":
		return png.Decode(img)
	case "image/jpeg":
		return jpeg.Decode(img)
	}

	return nil, errors.New("invalid format")
}

func (builder ImageBuilder) ResolveImageFile(filename string) (io.Reader, error) {
	if builder.IsImageURL(filename) {
		return builder.DownloadImageURL(filename)
	}

	path, err := filepath.Abs(filename)

	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	return file, err
}

func (builder ImageBuilder) Make(filename string) (image.Image, error) {
	img, err := builder.ResolveImageFile(filename)

	if err != nil {
		return nil, err
	}

	return builder.Decode(filename, img)
}

func (builder ImageBuilder) Insert(source image.Image, target image.Image, x int, y int) image.Image {
	offset := image.Pt(x, y)
	bounds := source.Bounds()
	output := image.NewRGBA(bounds)
	draw.Draw(output, bounds, source, image.ZP, draw.Src)
	draw.Draw(output, target.Bounds().Add(offset), target, image.ZP, draw.Over)

	return output
}

func (builder ImageBuilder) Save(source image.Image, filename string) error {
	output, err := os.Create(filename)

	if err != nil {
		return err
	}

	jpeg.Encode(output, source, &jpeg.Options{jpeg.DefaultQuality})

	defer output.Close()

	return nil
}
