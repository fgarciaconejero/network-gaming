package domain

type Repository interface {
	GenerateRandomNumber() int
	AddPoints(string, int) map[string]int
}
