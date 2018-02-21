package main

import (
	"testing"
)

type size_tests struct {
	version  int
	expected int
}

func table() []size_tests {
	tests := []size_tests{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}
	return tests
}

func TestVersionDeterminesSize(t *testing.T) {
	for _, test := range table() {
		size := Version(test.version).PatternSize()
		if size != test.expected {
			t.Errorf("Version %2d, expected %3d but got %3d", test.version,
				test.expected, size)
		}
	}
}

// func TestGenerateQRCodeReturnsValue(t *testing.T) {
// 	buffer := new(bytes.Buffer)
// 	GenerateQRCode(buffer, "555-2368")

// 	if buffer.Len() == 0 {
// 		t.Errorf("No QRCode generated")
// 	}
// 	_, err := png.Decode(buffer)

// 	if err != nil {
// 		t.Errorf("Generated QRCode is not a PNG: %s", err)
// 	}
// }
