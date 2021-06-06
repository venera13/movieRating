package domain

type RepositoryProvider interface {
	RatingRepository() RatingRepository
	MovieService() MovieService
}

type RatingUnitOfWork interface {
	RepositoryProvider
	UnitOfWork
}

type UnitOfWork interface {
	Complete(err *error)
}
