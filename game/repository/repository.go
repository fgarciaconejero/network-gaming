package repository

import (
	"math/rand"

	"github.com/fgarciaconejero/network-gaming/game/domain"
)

type GameRepository struct {
	pointStorage map[string]int
}

func NewGameRepository() domain.Repository {
	return &GameRepository{}
}

func (gr *GameRepository) AddPoints(id string, points int) map[string]int {
	gr.pointStorage[id] = points
	return gr.pointStorage
}

// Separated into another GameService function so that it can be mocked in tests
func (gr *GameRepository) GenerateRandomNumber() int {
	return rand.Intn(9) + 1
}
