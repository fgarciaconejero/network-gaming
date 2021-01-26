package repository

import (
	"math/rand"

	"github.com/fgarciaconejero/network-gaming/game/domain"
)

type GameRepository struct {
}

func NewGameRepository() domain.Repository {
	return &GameRepository{}
}

func (gr *GameRepository) AddPoints(id string, points int) error {
	return nil
}

// Separated into another GameService function so that it can be mocked in tests
func (gr *GameRepository) GenerateRandomNumber() int {
	return rand.Intn(9) + 1
}
