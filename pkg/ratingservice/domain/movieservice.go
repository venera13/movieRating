package domain

type MovieService interface {
	Get(id string) (*Movie, error)
}

type Movie struct {
	ID          string
	CreatedAt   int64
	UpdatedAt   int64
	DeletedAt   int64
	Name        string
	Description string
}
