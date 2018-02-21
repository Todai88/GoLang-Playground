package main

import (
	"image"
	"image/png"
	"io"
	"os"
)

func main() {
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	GenerateQRCode(file, "555-2368", Version(1))
}

// func GenerateQRCode(w io.Writer, code string) {
// 	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
// 	_ = png.Encode(w, img)
// }

func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := version.PatternSize()
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	return png.Encode(w, img)
}

type Version int8

func (v Version) PatternSize() int {
	return 4*int(v) + 17
}
