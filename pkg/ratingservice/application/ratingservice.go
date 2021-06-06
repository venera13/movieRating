package service

import (
	"github.com/google/uuid"
	"ratingservice/pkg/ratingservice/application/data"
	"ratingservice/pkg/ratingservice/application/errors"
	"ratingservice/pkg/ratingservice/domain"
	adapter "ratingservice/pkg/ratingservice/domain/adapter/movieservice"
)

type RatingService interface {
	RateTheMovie(request *data.RateTheMovieInput) error
}

type ratingService struct {
	ratingRepository domain.RatingRepository
	movieAdapter     adapter.MovieAdapter
}

func NewRatingService(
	ratingRepo domain.RatingRepository,
	movieAdapter adapter.MovieAdapter,
) RatingService {
	return &ratingService{
		ratingRepository: ratingRepo,
		movieAdapter:     movieAdapter,
	}
}

func (m *ratingService) RateTheMovie(request *data.RateTheMovieInput) error {
	if len(request.MovieId) == 0 {
		return errors.RequiredMovieIdError
	}

	if len(request.RatingValue) == 0 {
		return errors.RequiredRatingValueError
	}

	movie, err := m.movieAdapter.Get(request.MovieId)

	if movie == nil {
		return errors.MovieNotFound
	}

	ratingID := uuid.NewString()
	ratingData := domain.Rating{
		ID:          ratingID,
		MovieID:     request.MovieId,
		RatingValue: "5",
	}

	err = m.ratingRepository.AddRating(ratingData)

	return err
}
