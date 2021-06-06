package infrastructure

import (
	"database/sql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"ratingservice/pkg/ratingservice/domain"
)

type unitOfWorkFactory struct {
	client sql.DB
}

type unitOfWork struct {
	transaction Transaction
}

func (u *unitOfWorkFactory) NewUnitOfWork() (domain.RatingUnitOfWork, error) {
	log.Info("NewUnitOfWork start")
	transaction, err := u.client.Begin()
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{
		"NewUnitOfWork exit": transaction,
	}).Info("NewUnitOfWork exit")
	return &unitOfWork{transaction: transaction}, nil
}

func (u *unitOfWork) RatingRepository() domain.RatingRepository {
	return &DatabaseRepository{transaction: u.transaction}
}

func (u *unitOfWork) MovieAdapter() domain.MovieAdapter {
	return &MovieRepository{transaction: u.transaction}
}

func (u *unitOfWork) Complete(err *error) error {
	if err != nil {
		err2 := u.transaction.Rollback()
		if err2 != nil {
			return errors.Wrap(*err, err2.Error())
		}
		return *err
	}
	return errors.WithStack(u.transaction.Commit())
}
