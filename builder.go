package ImageBuilder

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"net/http"
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
