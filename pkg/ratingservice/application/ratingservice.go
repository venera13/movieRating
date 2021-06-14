package service

import (
	"github.com/google/uuid"
	"ratingservice/pkg/ratingservice/application/data"
	"ratingservice/pkg/ratingservice/application/errors"
	"ratingservice/pkg/ratingservice/application/provider"
	"ratingservice/pkg/ratingservice/application/unitofwork"
	"ratingservice/pkg/ratingservice/domain"
)

type RatingService struct {
	unitOfWorkFactory unitofwork.UnitOfWorkFactory
	movieProvider     provider.MovieProvider
}

func NewRatingService(
	unitOfWorkFactory unitofwork.UnitOfWorkFactory,
	movieProvider provider.MovieProvider,
) RatingService {
	return RatingService{
		unitOfWorkFactory: unitOfWorkFactory,
		movieProvider:     movieProvider,
	}
}

func (srv *RatingService) RateTheMovie(request *data.RateTheMovieInput) error {
	if len(request.MovieId) == 0 {
		return errors.RequiredMovieIdError
	}

	if len(request.RatingValue) == 0 {
		return errors.RequiredRatingValueError
	}

	movie, err := srv.movieProvider.Get(request.MovieId)

	if movie == nil {
		return errors.MovieNotFound
	}

	var unitOfWork unitofwork.RatingUnitOfWork
	unitOfWork, err = srv.unitOfWorkFactory.NewUnitOfWork()
	if err != nil {
		return err
	}
	defer func() {
		unitOfWork.Complete(&err)
	}()

	ratingService := unitOfWork.RatingRepository()

	ratingID := uuid.NewString()
	ratingData := domain.Rating{
		ID:          ratingID,
		MovieID:     request.MovieId,
		RatingValue: request.RatingValue,
	}

	err = ratingService.AddRating(ratingData)

	return err
}
