package model

type Game struct {
	players []Player
}

type Player struct {
	FirstNumber  int
	SecondNumber int
}
