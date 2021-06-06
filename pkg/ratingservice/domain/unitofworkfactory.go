package domain

type UnitOfWorkFactory interface {
	NewUnitOfWork() (RatingUnitOfWork, error)
}
