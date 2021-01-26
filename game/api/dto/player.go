package dto

import "github.com/fgarciaconejero/network-gaming/game/domain/model"

type Player struct {
	FirstNumber  int `json:"first_value" validate:"required"`
	SecondNumber int `json:"second_value" validate:"required"`
}

func (p *Player) ToModel() *model.Player {
	return &model.Player{
		FirstNumber:  p.FirstNumber,
		SecondNumber: p.SecondNumber,
	}
}
