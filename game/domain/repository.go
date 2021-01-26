package domain

type Repository interface {
	AddPoints(string, int) error
}
