package service

import (
	"errors"
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
		return serviceerrors.RequiredMovieIdError
	}

	if request.RatingValue == 0 {
		return serviceerrors.RequiredRatingValueError
	}

	movie, err := srv.movieProvider.Get(request.MovieId)

	if movie == nil {
		return serviceerrors.MovieNotFound
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

	var movieRating *domain.Rating
	movieRating, err = getRatingByMovieId(ratingService, movie.ID)

	if errors.Is(err, serviceerrors.MovieNotFound) {
		err = addNewRating(ratingService, request.MovieId, request.RatingValue)
	}

	if err != nil {
		return err
	}

	err = addRating(ratingService, movieRating, request.RatingValue)

	return err
}

func getRatingByMovieId(repository domain.RatingRepository, movieId string) (*domain.Rating, error) {
	rating, err := repository.GetRatingByMovieId(movieId)

	if errors.Is(err, domain.ErrorMovieNotFound) {
		return rating, serviceerrors.MovieNotFound
	}

	if err != nil {
		return rating, err
	}

	return rating, nil
}

func addNewRating(repository domain.RatingRepository, movieId string, ratingValue int64) error {
	ratingID := uuid.NewString()
	ratingData := domain.Rating{
		ID:            ratingID,
		MovieID:       movieId,
		RatingValue:   ratingValue,
		NumberRatings: 1,
	}

	err := repository.Add(ratingData)

	return err
}

func addRating(repository domain.RatingRepository, movieRating *domain.Rating, ratingValue int64) error {
	movieRating.NumberRatings = movieRating.NumberRatings + 1
	movieRating.RatingValue = calcRating(movieRating, ratingValue)

	err := repository.Update(*movieRating)

	return err
}

func calcRating(movieRating *domain.Rating, ratingValue int64) int64 {
	return (movieRating.RatingValue + ratingValue) / movieRating.NumberRatings
}
