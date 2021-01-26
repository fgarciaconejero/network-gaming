package domain

import "github.com/gin-gonic/gin"

type API interface {
	Start(g *gin.Context)
}
