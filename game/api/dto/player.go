package dto

import "../domain/model"

type Player struct {
	FirstValue  int `json:"first_value" validate:"required"`
	SecondValue int `json:"second_value" validate:"required"`
}

func (p *Player) ToModel() *model.Player {
	return &model.Player{
		FirstValue:  p.FirstValue,
		SecondValue: p.SecondValue,
	}
}
