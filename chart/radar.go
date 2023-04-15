package chart

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/chromedp/chromedp"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// static data for dev
var radarDataBJ = [][]float32{
	{80, 6, 21, 11, 58, 17, 17},
}

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
	{Name: "Git", Max: 100},
	{Name: "Algorithm", Max: 100},
	{Name: "Soft Skills", Max: 60},
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

type RadarExamples struct{}

func (RadarExamples) Generate() {
	page := components.NewPage()
	page.AddCharts(
		radarStyle(),
	)
	f, err := os.Create("radar.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the HTML page
	var buf []byte
	// get full path
	_, filename, _, _ := runtime.Caller(0)
	filename = filename[:len(filename)-len("chart/radar.go")]
	log.Println("file://" + filename)
	if err := chromedp.Run(ctx, chromedp.Navigate("file://"+filename+f.Name()),
		chromedp.CaptureScreenshot(&buf),
	); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("radar.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}
}
