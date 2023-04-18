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
			Level:             40,
			Skills: [][]float32{
				{110, 15, 2},
			},
		},
	)
}
