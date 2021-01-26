package domain

import (
	"../domain/model"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Start(g *gin.Context, players []model.Player) error
}
