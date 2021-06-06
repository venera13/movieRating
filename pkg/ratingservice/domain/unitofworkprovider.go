package domain

type RepositoryProvider interface {
	RatingRepository() RatingRepository
	MovieAdapter() MovieAdapter
}

type RatingUnitOfWork interface {
	RepositoryProvider
	UnitOfWork
}

type UnitOfWork interface {
	Complete(err *error)
}
