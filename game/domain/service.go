package domain

import (
	"github.com/fgarciaconejero/network-gaming/game/domain/model"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Start(g *gin.Context, players []model.Player) (bool, error)
}
