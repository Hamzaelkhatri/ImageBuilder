package ImageBuilder

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"os"

	"github.com/Hamzaelkhatri/ImageBuilder/v2/chart"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

func Builder(arg []string) string {
	imageURL := arg
	outImg := image.NewRGBA(image.Rect(0, 0, 600, 600))

	x := 0
	y := 0
	for _, url := range imageURL {
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		srcImg, _, err := image.Decode(resp.Body)
		if err != nil {
			// panic(err)
			fmt.Println("Error: ", err)
			return "Error"
		}

		draw.Draw(outImg, image.Rect(x, y, x+300, y+300), srcImg, image.Point{0, 0}, draw.Src)
		x += 300
		if x == 600 {
			x = 0
			y += 300
		}
	}
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err := png.Encode(w, outImg)
	if err != nil {
		panic(err)
	}
	w.Flush()
	s := base64.StdEncoding.EncodeToString(buf.Bytes())
	return s
}

func CardProfile(Level int, Login string, Avatar string) string {
	Card := image.NewRGBA(image.Rect(0, 0, 1500, 1000))
	e := chart.RadarExamples{}
	e.Generate()
	// Background color
	draw.Draw(Card, image.Rect(0, 0, 1500, 1000), image.NewUniform(color.RGBA{22, 22, 39, 255}), image.Point{0, 0}, draw.Src)

	// Profile picture
	resp, err := http.Get(Avatar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	srcImg, _, err := image.Decode(resp.Body)
	if err != nil {
		panic(err)
	}
	draw.Draw(Card, image.Rect(0, 0, 1000, 1000), srcImg, image.Point{0, 0}, draw.Src)
	// get image from local file
	file, err := os.Open("radar.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	srcImg, _, err = image.Decode(file)
	if err != nil {
		panic(err)
	}
	// crop image
	srcImg = srcImg.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(100, 10, 700, 500))
	// add radar chart
	draw.Draw(Card, image.Rect(700, 0, 1500, 1000), srcImg, image.Point{0, 0}, draw.Src)

	// add text
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size:    72,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	d := &font.Drawer{
		Dst:  Card,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
		Dot:  fixed.P(400, 100),
	}

	d.DrawString(Login)
	d.Dot = fixed.P(400, 200)
	d.DrawString(fmt.Sprintf("Level: %d", Level))

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err = png.Encode(w, Card)
	if err != nil {
		panic(err)
	}
	w.Flush()
	// save to file
	file, err = os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, Card)
	if err != nil {
		panic(err)
	}
	return "Done"
}
