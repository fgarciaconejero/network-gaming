package internal

import (
	"github.com/fgarciaconejero/network-gaming/game/domain"
)

func (r *SRV) AddHandlers(gh domain.API) *SRV {
	r = AddPingHandler(r)
	r = AddGameHandlers(r, gh)
	return r
}

func AddPingHandler(r *SRV) *SRV {
	healthCheckHandler := &HealthChecker{}

	route := r.Group("/ping")
	route.GET("", healthCheckHandler.Get)
	return r
}

func AddGameHandlers(r *SRV, gameHandler domain.API) *SRV {
	route := r.Group("game/")
	route.POST("start", gameHandler.Start)
	return r
}
