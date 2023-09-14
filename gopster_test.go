package gopster

import (
	"image/jpeg"
	"os"
	"testing"
)

func TestGopster(t *testing.T) {
	c, err := NewChart(
		ShowTitles(),
		Width(3),
		Height(3),
	)
	if err != nil {
		t.Fatalf("Error creating new chart: %s", err.Error())
	}

	f, err := os.Open("testdata/bomb.jpg")
	if err != nil {
		t.Fatalf("Error opening test image: %s", err.Error())
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		t.Fatalf("Error decoding test image: %s", err.Error())
	}

	for i := 0; i < 9; i++ {
		c.MustAddItem("Get Warmer", "Bomb the Music Industry!", img)
	}

	newImg := c.Generate()
	out, err := os.Create("out.jpg")
	if err != nil {
		t.Fatalf("Error creating new file: %s", err.Error())
	}

	if err := jpeg.Encode(out, newImg, nil); err != nil {
		t.Fatalf("Error encoding new image: %s", err.Error())
	}
}
