package main

import (
	"github.com/Hamzaelkhatri/ImageBuilder/v2"
)

func main() {
	ImageBuilder.Init(
		ImageBuilder.CardData{
			Name:              "Hamza",
			NumberOfExercises: 110,
			Avatar:            "https://learn.reboot01.com/git/avatars/9870e141f7a57c7d0b3e082d9cf97219?size=40",
			Level:             29,
			Raids: []ImageBuilder.Raid{
				{
					Name:   "Quad",
					Status: "done",
					Grade:  1,
				},
				{
					Name:   "Sudoku",
					Status: "done",
					Grade:  0,
				},
			},
			Skills: [][]float32{
				{60, 50, 25, 18, 20, 10, 30},
			},
		},
	)
}
