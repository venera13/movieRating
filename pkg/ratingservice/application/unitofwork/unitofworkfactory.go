package serviceunitofwork

type UnitOfWorkFactory interface {
	NewUnitOfWork() (RatingUnitOfWork, error)
}
