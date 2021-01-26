package internal

import (
	"github.com/fgarciaconejero/network-gaming/game/api"
	"github.com/fgarciaconejero/network-gaming/game/domain"
)

func (r *SRV) AddHandlers() *SRV {
	r = AddGameHandlers(r, api.NewGameHandler())
	r = AddPingHandler(r)
	return r
}

func AddPingHandler(r *SRV) *SRV {
	healthCheckHandler := &HealthChecker{}

	route := r.Group("/ping")
	route.Use()
	route.GET("", healthCheckHandler.Get)
	return r
}

func AddGameHandlers(r *SRV, gameHandler domain.API) *SRV {
	route := r.Group("game/")
	route.POST("start", gameHandler.Start)
	return r
}
