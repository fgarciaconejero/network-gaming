package dto

import "github.com/fgarciaconejero/network-gaming/game/domain/model"

type Player struct {
	ID           string `json:"id" validate:"required"`
	FirstNumber  int    `json:"first_number" validate:"required"`
	SecondNumber int    `json:"second_number" validate:"required"`
}

func (p *Player) ToModel() *model.Player {
	return &model.Player{
		ID:           p.ID,
		FirstNumber:  p.FirstNumber,
		SecondNumber: p.SecondNumber,
	}
}

func (c *Player) FromModel(pm *model.Player) {
	c.ID = pm.ID
	c.FirstNumber = pm.FirstNumber
	c.SecondNumber = pm.SecondNumber
}
