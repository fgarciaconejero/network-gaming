package main

import (
	"./internal"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.Use(gin.Logger())
	internal.NewServer(g, "8080").AddHandlers().Run()
}
