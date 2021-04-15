package lib

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/nfnt/resize"
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

func (builder ImageBuilder) hexToRGBA(hex string) (color.RGBA, error) {
	var (
		rgba             color.RGBA
		err              error
		errInvalidFormat = fmt.Errorf("invalid")
	)
	rgba.A = 0xff
	if hex[0] != '#' {
		return rgba, errInvalidFormat
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}
	switch len(hex) {
	case 7:
		rgba.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		rgba.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		rgba.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		rgba.R = hexToByte(hex[1]) * 17
		rgba.G = hexToByte(hex[2]) * 17
		rgba.B = hexToByte(hex[3]) * 17
	default:
		err = errInvalidFormat
	}
	return rgba, err
}

func (builder ImageBuilder) Create(width int, height int, color string) (image.Image, error) {
	bgColor, err := builder.hexToRGBA(color)

	if err != nil {
		return nil, err
	}

	background := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(background, background.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	return background, err
}

func (builder ImageBuilder) Resize(img image.Image, width uint, height uint) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
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
