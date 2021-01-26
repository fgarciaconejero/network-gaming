package repository

import "github.com/fgarciaconejero/network-gaming/game/domain"

type GameRepository struct {
}

func NewGameRepository() domain.Repository {
	return &GameRepository{}
}

func (gr *GameRepository) AddPoints(id string, points int) error {
	return nil
}
