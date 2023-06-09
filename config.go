package ImageBuilder

import (
	"github.com/Hamzaelkhatri/ImageBuilder/v2/chart"
)

type Raid struct {
	Name   string
	Status string
	Grade  float32
}

type Checkpoint struct {
	Name     string
	Status   string
	Grade    float32
	MaxLevel int
}

type CardData struct {
	Name              string
	NumberOfExercises int
	Avatar            string
	Level             float64
	Raids             []Raid
	Checkpoints       []Checkpoint
	Skills            [][]float32
}

func Init(card CardData) string {
	char := chart.Radar{}
	bas := char.Generate(card.Skills)
	return CardProfile(card, bas)
}
