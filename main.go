package main

import (
	"github.com/fgarciaconejero/network-gaming/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.Use(gin.Logger())
	internal.NewServer(g, "8080").AddHandlers().Run()
}
