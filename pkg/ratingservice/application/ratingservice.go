package service

import (
	"github.com/google/uuid"
	"ratingservice/pkg/ratingservice/application/data"
	"ratingservice/pkg/ratingservice/application/errors"
	"ratingservice/pkg/ratingservice/domain"
)

type RatingService struct {
	unitOfWorkFactory domain.UnitOfWorkFactory
}

func NewRatingService(
	unitOfWorkFactory domain.UnitOfWorkFactory,
) RatingService {
	return RatingService{
		unitOfWorkFactory: unitOfWorkFactory,
	}
}

func (srv *RatingService) RateTheMovie(request *data.RateTheMovieInput) error {
	if len(request.MovieId) == 0 {
		return errors.RequiredMovieIdError
	}

	if len(request.RatingValue) == 0 {
		return errors.RequiredRatingValueError
	}

	unitOfWork, err := srv.unitOfWorkFactory.NewUnitOfWork()
	if err != nil {
		return err
	}
	defer func() {
		unitOfWork.Complete(&err)
	}()

	movieService := unitOfWork.MovieService()
	ratingService := unitOfWork.RatingRepository()

	var movie *domain.Movie
	movie, err = movieService.Get(request.MovieId)

	if movie == nil {
		return errors.MovieNotFound
	}

	ratingID := uuid.NewString()
	ratingData := domain.Rating{
		ID:          ratingID,
		MovieID:     request.MovieId,
		RatingValue: request.RatingValue,
	}

	err = ratingService.AddRating(ratingData)

	return err
}
