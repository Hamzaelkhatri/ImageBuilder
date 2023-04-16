package ImageBuilder

import "github.com/Hamzaelkhatri/ImageBuilder/v2/chart"

type Raid struct {
	Name   string
	Status string
	Grade  float32
}

type CardData struct {
	Name              string
	NumberOfExercises int
	Avatar            string
	Level             int
	Raids             []Raid
	Skills            [][]float32
}

func Init(card CardData) {
	char := chart.Radar{}
	bas := char.Generate(card.Skills)
	e := CardProfile(card, bas)
	println(e)
}