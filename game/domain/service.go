package domain

import (
	"context"

	"github.com/fgarciaconejero/network-gaming/game/domain/model"
)

type Service interface {
	Start(g context.Context, players []model.Player) error
}
