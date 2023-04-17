package chart

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// static data for dev
var radarDataBJ = [][]float32{}

func generateRadarItems(radarData [][]float32) []opts.RadarData {
	items := make([]opts.RadarData, 0)
	for i := 0; i < len(radarData); i++ {
		items = append(items, opts.RadarData{Value: radarData[i]})
	}
	return items
}

var indicators = []*opts.Indicator{
	{Name: "Golang", Max: 100},
	{Name: "Math", Max: 100},
	{Name: "Problem Solving", Max: 100},
	{Name: "Unix", Max: 100},
	{Name: "Git", Max: 50},
	{Name: "Algorithm", Max: 100},
	{Name: "Soft Skills", Max: 70},
}

func radarStyle() *charts.Radar {
	radar := charts.NewRadar()
	radar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:      "Skills",
			Right:      "center",
			TitleStyle: &opts.TextStyle{Color: "#eee"},
		}),
		charts.WithInitializationOpts(opts.Initialization{
			BackgroundColor: "#161627",
		}),
		charts.WithRadarComponentOpts(opts.RadarComponent{
			Indicator:   indicators,
			Shape:       "circle",
			SplitNumber: 5,
			SplitLine: &opts.SplitLine{
				Show: true,
				LineStyle: &opts.LineStyle{
					Opacity: 0.1,
				},
			},
		}),

		charts.WithLegendOpts(opts.Legend{
			Show:   true,
			Bottom: "5px",
			TextStyle: &opts.TextStyle{
				Color: "#eee",
			},
		}),
	)

	radar.AddSeries("", generateRadarItems(radarDataBJ)).
		SetSeriesOptions(
			charts.WithLineStyleOpts(opts.LineStyle{
				Width:   1,
				Opacity: 0.5,
			}),
			charts.WithAreaStyleOpts(opts.AreaStyle{
				Opacity: 0.1,
			}),
			charts.WithItemStyleOpts(opts.ItemStyle{
				Color: "#F9713C",
			}),
		)

	return radar
}

type Radar struct{}

func (Radar) Generate(Data [][]float32) string {
	radarDataBJ = Data
	page := components.NewPage()
	page.AddCharts(
		radarStyle(),
	)
	f, err := os.CreateTemp("", "radar*.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the HTML page
	var buf []byte
	// get full of temporary of radar.html
	path := f.Name()
	// _, filename
	// filename = filename[:len(filename)-len("chart/radar.go")]
	log.Println("Tmporary ", path)
	if err := chromedp.Run(ctx, chromedp.Navigate("file://"+path),
		chromedp.CaptureScreenshot(&buf),
	); err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(buf)
}
