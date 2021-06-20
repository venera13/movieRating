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
	unitOfWorkFactory serviceunitofwork.UnitOfWorkFactory
	movieProvider     provider.MovieProvider
}

func NewRatingService(
	unitOfWorkFactory serviceunitofwork.UnitOfWorkFactory,
	movieProvider provider.MovieProvider,
) RatingService {
	return RatingService{
		unitOfWorkFactory: unitOfWorkFactory,
		movieProvider:     movieProvider,
	}
}

func (srv *RatingService) RateTheMovie(request *data.RateTheMovieInput) error {
	if request.MovieID == "" {
		return serviceerrors.ErrorRequiredMovieId
	}

	if request.RatingValue == 0 {
		return serviceerrors.ErrorRequiredRatingValue
	}

	movie, err := srv.movieProvider.Get(request.MovieID)

	if movie == nil {
		return serviceerrors.ErrorMovieNotFound
	}

	var unitOfWork serviceunitofwork.RatingUnitOfWork
	unitOfWork, err = srv.unitOfWorkFactory.NewUnitOfWork()

	if err != nil {
		return err
	}

	defer func() {
		unitOfWork.Complete(&err)
	}()

	ratingService := unitOfWork.RatingRepository()

	var movieRating *domain.Rating
	movieRating, err = getRatingByMovieID(ratingService, movie.ID)

	if errors.Is(err, serviceerrors.ErrorMovieNotFound) {
		err = addNewRating(ratingService, request.MovieID, request.RatingValue)
	}

	if err != nil {
		return err
	}

	err = addRating(ratingService, movieRating, request.RatingValue)

	return err
}

func getRatingByMovieID(repository domain.RatingRepository, movieID string) (*domain.Rating, error) {
	rating, err := repository.GetRatingByMovieID(movieID)

	if errors.Is(err, domain.ErrorMovieNotFound) {
		return rating, serviceerrors.ErrorMovieNotFound
	}

	if err != nil {
		return rating, err
	}

	return rating, nil
}

func addNewRating(repository domain.RatingRepository, movieID string, ratingValue int64) error {
	ratingID := uuid.NewString()
	ratingData := domain.Rating{
		ID:            ratingID,
		MovieID:       movieID,
		RatingValue:   ratingValue,
		NumberRatings: 1,
		DeletedAt:     "nil",
	}

	err := repository.Add(ratingData)

	return err
}

func addRating(repository domain.RatingRepository, movieRating *domain.Rating, ratingValue int64) error {
	movieRating.NumberRatings++
	movieRating.RatingValue = calcRating(movieRating, ratingValue)

	err := repository.Update(*movieRating)

	return err
}

func calcRating(movieRating *domain.Rating, ratingValue int64) int64 {
	return (movieRating.RatingValue + ratingValue) / movieRating.NumberRatings
}
