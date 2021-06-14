package unitofwork

import (
	"ratingservice/pkg/ratingservice/domain"
)

type RepositoryProvider interface {
	RatingRepository() domain.RatingRepository
}

type RatingUnitOfWork interface {
	RepositoryProvider
	UnitOfWork
}

type UnitOfWork interface {
	Complete(err *error)
}
