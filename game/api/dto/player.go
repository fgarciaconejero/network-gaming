package dto

import "github.com/fgarciaconejero/network-gaming/game/domain/model"

type Player struct {
	FirstNumber  int `json:"first_number" validate:"required"`
	SecondNumber int `json:"second_number" validate:"required"`
}

func (p *Player) ToModel() *model.Player {
	return &model.Player{
		FirstNumber:  p.FirstNumber,
		SecondNumber: p.SecondNumber,
	}
}

func (c *Player) FromModel(pm *model.Player) {
	c.FirstNumber = pm.FirstNumber
	c.SecondNumber = pm.SecondNumber
}
