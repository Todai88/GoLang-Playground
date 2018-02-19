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

	GenerateQRCode(file, "555-2368")
}

func GenerateQRCode(w io.Writer, code string) {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	_ = png.Encode(w, img)
}
