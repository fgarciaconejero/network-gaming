package service

import (
	"context"

	"github.com/fgarciaconejero/network-gaming/game/domain"
	"github.com/fgarciaconejero/network-gaming/game/domain/model"
	"github.com/fgarciaconejero/network-gaming/game/repository"
)

type GameService struct {
	GameRepository domain.Repository
}

func NewGameService() domain.Service {
	gr := repository.NewGameRepository()
	return &GameService{GameRepository: gr}
}

func (gs *GameService) Start(g context.Context, players []model.Player) error {
	serverNumber := gs.GameRepository.GenerateRandomNumber()
	for _, v := range players {
		if v.FirstNumber == serverNumber || v.SecondNumber == serverNumber {
			gs.GameRepository.AddPoints(v.ID, 5)
		} else if serverNumber > v.FirstNumber && serverNumber < v.SecondNumber {
			aux := 5 - (v.SecondNumber - v.FirstNumber)
			if aux > 0 {
				gs.GameRepository.AddPoints(v.ID, 5-(v.SecondNumber-v.FirstNumber))
			}
		} else {
			gs.GameRepository.AddPoints(v.ID, -1)
		}
	}
	return nil
}
