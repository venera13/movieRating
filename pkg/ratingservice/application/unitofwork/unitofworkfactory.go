package unitofwork

type UnitOfWorkFactory interface {
	NewUnitOfWork() (RatingUnitOfWork, error)
}
