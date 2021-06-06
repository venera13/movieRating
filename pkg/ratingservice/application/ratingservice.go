package service

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"ratingservice/pkg/ratingservice/application/data"
	"ratingservice/pkg/ratingservice/application/errors"
	"ratingservice/pkg/ratingservice/domain"
)

type RatingService struct {
	unitOfWorkFactory domain.UnitOfWorkFactory
}

//func NewRatingService(
//	unitOfWorkFactory domain.UnitOfWorkFactory,
//	movieAdapter adapter.MovieAdapter,
//) RatingService {
//	return &ratingService{
//		ratingRepository: ratingRepo,
//		movieAdapter:     movieAdapter,
//	}

func (srv *RatingService) RateTheMovie(request *data.RateTheMovieInput) error {
	if len(request.MovieId) == 0 {
		return errors.RequiredMovieIdError
	}

	if len(request.RatingValue) == 0 {
		return errors.RequiredRatingValueError
	}

	log.Info("сейчас перейдет в NewUnitOfWork")
	log.WithFields(log.Fields{
		"NewUnitOfWork": srv.unitOfWorkFactory,
	}).Info("NewUnitOfWork exit")
	unitOfWork, err := srv.unitOfWorkFactory.NewUnitOfWork()
	if err != nil {
		return err
	}
	defer func() {
		err = unitOfWork.Complete(&err)
	}()

	ratingService := unitOfWork.RatingRepository()
	movieAdapter := unitOfWork.MovieAdapter()

	movie, err := movieAdapter.Get(request.MovieId)

	if movie == nil {
		return errors.MovieNotFound
	}

	ratingID := uuid.NewString()
	ratingData := domain.Rating{
		ID:          ratingID,
		MovieID:     request.MovieId,
		RatingValue: "5",
	}

	err = ratingService.AddRating(ratingData)

	return err
}
