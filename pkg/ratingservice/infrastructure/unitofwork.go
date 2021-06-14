package infrastructure

import (
	"database/sql"
	"ratingservice/pkg/ratingservice/application/unitofwork"
	"ratingservice/pkg/ratingservice/domain"
)

func CreateUnitOfWorkFactory(db *sql.DB) unitofwork.UnitOfWorkFactory {
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

func (u *UnitOfWorkFactory) NewUnitOfWork() (unitofwork.RatingUnitOfWork, error) {
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
	if *err != nil {
		err2 := u.transaction.Rollback()
		err = &err2
	} else {
		err2 := u.transaction.Commit()
		err = &err2
	}
}
