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
			Level:             30.55,
			Raids: []ImageBuilder.Raid{
				{Name: "sudoku", Status: "Absent", Grade: 1},
				{Name: "quadchecker", Status: "done", Grade: 0},
				{Name: "Checkpoint 03", Status: "done", Grade: 1},
				{Name: "Checkpoint 02", Status: "done", Grade: 0.75},
				{Name: "Checkpoint 01", Status: "done", Grade: 1},
				{Name: "quad", Status: "done", Grade: 1},
			},
			Skills: [][]float32{
				{110, 15, 2},
			},
		},
	)
}
