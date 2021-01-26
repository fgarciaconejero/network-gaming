package main

import (
	"github.com/fgarciaconejero/network-gaming/game/api"
	"github.com/fgarciaconejero/network-gaming/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.Use(gin.Logger())
	auxApi := api.NewGameHandler()
	internal.NewServer(g, "8080").AddHandlers(auxApi).Run()
}
