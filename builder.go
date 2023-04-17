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

	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/nfnt/resize"
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

func drawRectangle(img draw.Image, x, y, w, h, rx, ry int, c color.Color) {
	ctx := draw2dimg.NewGraphicContext(img)
	ctx.SetFillColor(c)
	draw2dkit.RoundedRectangle(ctx, float64(x), float64(y), float64(w), float64(h), float64(rx), float64(ry))
	ctx.FillStroke()
}

func CardProfile(card CardData, radar string) string {
	Card := image.NewRGBA(image.Rect(0, 0, 1500, 1000))
	draw.Draw(Card, image.Rect(0, 0, 1500, 1000), image.NewUniform(color.RGBA{22, 22, 39, 255}), image.Point{0, 0}, draw.Src)

	srcImg, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(radar)))
	if err != nil {
		panic(err)
	}
	// crop image
	srcImg = srcImg.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(100, 10, 700, 500))
	// resize image
	srcImg = resize.Resize(800, 0, srcImg, resize.Lanczos3)

	// add text
	f, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size:    60,
		DPI:     60,
		Hinting: font.Hinting(font.StretchCondensed),
	})
	d := &font.Drawer{
		Dst:  Card,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
		Dot:  fixed.P(600, 100),
	}

	d.DrawString(fmt.Sprintf("%s's Stats", card.Name))
	percent := float64((float64(card.Level) / 40) * 100)
	drawRectangle(Card, 100, 500+400, 700+700, 820, 60, -60, color.RGBA{0, 0, 0, 50})
	drawRectangle(Card, 100, 500+400, (400-(int(percent)))+(int(percent)*10), 820, 60, -60, color.RGBA{32, 200, 93, 255})

	d.Dot = fixed.P(600, 880)
	d.DrawString(fmt.Sprintf("level %d - %.2f %%", card.Level, percent))

	face = truetype.NewFace(f, &truetype.Options{
		Size:    45,
		DPI:     45,
		Hinting: font.Hinting(font.StretchCondensed),
	})

	d = &font.Drawer{
		Dst:  Card,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
		Dot:  fixed.P(100, 800),
	}

	d.DrawString("Piscine Level")

	d.Dot = fixed.P(230, 290)
	d.DrawString("Piscine Stats")

	d.Dot = fixed.P(70, 350)
	face = truetype.NewFace(f, &truetype.Options{
		Size:    45,
		DPI:     40,
		Hinting: font.Hinting(font.StretchCondensed),
	})
	d.Face = face
	d.DrawString("Exercises done")

	d.Dot = fixed.P(400, 350)
	d.Src = image.NewUniform(color.RGBA{86, 186, 16, 255})
	d.DrawString(fmt.Sprintf("%d ", card.NumberOfExercises))

	d.Src = image.NewUniform(color.RGBA{255, 255, 255, 255})
	d.Face = face
	d.DrawString("/ 122")
	for i, raid := range card.Raids {
		d.Src = image.NewUniform(color.RGBA{255, 255, 255, 255})
		d.Dot = fixed.P(70, 400+(i*50))
		d.DrawString(raid.Name)
		d.Dot = fixed.P(400, 400+(i*50))
		if raid.Status != "done" {
			d.Src = image.NewUniform(color.RGBA{241, 210, 54, 255})
			d.DrawString(raid.Status)
			continue
		}
		if raid.Grade >= 1 {
			d.Src = image.NewUniform(color.RGBA{86, 186, 16, 255})
			d.DrawString("Succeeded")
		} else {
			d.Src = image.NewUniform(color.RGBA{255, 0, 0, 255})
			d.DrawString("Failed")
		}
		// d.DrawString(fmt.Sprintf(" %s", raid.Status))
	}

	// draw a rectangle
	ctx := draw2dimg.NewGraphicContext(Card)

	// Rectangle
	ctx.BeginPath()      // Initialize a new path
	ctx.MoveTo(600, 650) // Move to (600, 600)
	ctx.LineTo(600, 250)
	ctx.LineTo(50, 250)
	ctx.LineTo(50, 650)
	ctx.Close() // Close the path
	// add stroke
	ctx.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	ctx.SetLineWidth(5)
	ctx.Stroke()
	// add transparent fill
	ctx.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0x44})

	ctx.FillStroke()

	draw.Draw(Card, image.Rect(600, 150, 1500, 1000), srcImg, image.Point{0, 0}, draw.Src)

	logo, err := http.Get(card.Avatar)
	if err != nil {
		panic(err)
	}
	defer logo.Body.Close()
	srcImg, _, err = image.Decode(logo.Body)
	if err != nil {
		panic(err)
	}
	// add logo

	srcImg = resize.Resize(150, 0, srcImg, resize.NearestNeighbor)

	draw.Draw(Card, image.Rect(50, 25, 1500, 1000), srcImg, image.Point{0, 0}, draw.Src)

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	err = png.Encode(w, Card)
	if err != nil {
		panic(err)
	}
	w.Flush()
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
