package infrastructure

import (
	"database/sql"
	"github.com/pkg/errors"
	"ratingservice/pkg/ratingservice/application/unitofwork"
	"ratingservice/pkg/ratingservice/domain"
)

func CreateUnitOfWorkFactory(db *sql.DB) serviceunitofwork.UnitOfWorkFactory {
	return &UnitOfWorkFactory{
		client: db,
	}
}

type UnitOfWorkFactory struct {
	client *sql.DB
}

type unitOfWork struct {
	transaction Transaction
}

func (u *UnitOfWorkFactory) NewUnitOfWork() (serviceunitofwork.RatingUnitOfWork, error) {
	transaction, err := u.client.Begin()

	if err != nil {
		return nil, err
	}

	return &unitOfWork{transaction: transaction}, nil
}

func (u *unitOfWork) RatingRepository() domain.RatingRepository {
	return &DatabaseRepository{transaction: u.transaction}
}

func (u *unitOfWork) Complete(err *error) {
	var err2 error
	if *err != nil {
		err2 = u.transaction.Rollback()
	} else {
		err2 = u.transaction.Commit()
	}

	if err2 != nil {
		err2 = errors.Wrap(*err, err2.Error())
		err = &err2
	}
}
