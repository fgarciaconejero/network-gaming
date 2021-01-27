package domain

type Repository interface {
	GenerateRandomNumber() int
	AddPoints(string, int)
	GetPoints() map[string]int
}
