package service

import (
	"github.com/fgarciaconejero/network-gaming/game/domain"
	"github.com/fgarciaconejero/network-gaming/game/domain/model"
	"github.com/gin-gonic/gin"
)

type GameService struct {
	GameRepository domain.Repository
}

func NewGameService() domain.Service {
	return &GameService{}
}

func (gs *GameService) Start(g *gin.Context, players []model.Player) error {

	return nil
}
