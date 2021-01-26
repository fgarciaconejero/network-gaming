package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthChecker struct{}

func (h HealthChecker) Get(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
