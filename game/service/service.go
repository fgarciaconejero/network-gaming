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

func (gs *GameService) Start(g context.Context, players []model.Player) string {
	serverNumber := gs.GameRepository.GenerateRandomNumber()
	var points map[string]int
	for _, v := range players {
		if v.FirstNumber == serverNumber || v.SecondNumber == serverNumber {
			points = gs.GameRepository.AddPoints(v.ID, 5)
		} else if serverNumber > v.FirstNumber && serverNumber < v.SecondNumber {
			aux := 5 - (v.SecondNumber - v.FirstNumber)
			if aux > 0 {
				points = gs.GameRepository.AddPoints(v.ID, 5-(v.SecondNumber-v.FirstNumber))
			}
		} else {
			points = gs.GameRepository.AddPoints(v.ID, -1)
		}

		if points[v.ID] == 21 {
			return v.ID
		}
	}
	return ""
}
