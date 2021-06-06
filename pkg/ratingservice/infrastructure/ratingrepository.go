package infrastructure

import (
	"ratingservice/pkg/ratingservice/domain"
)

type DatabaseRepository struct {
	transaction Transaction
}

func (ratingRepo *DatabaseRepository) AddRating(ratingData domain.Rating) error {
	query := "INSERT INTO rating(id, movie_id, rating_value) VALUES (?, ?, ?)"
	_, err := ratingRepo.transaction.Exec(query, ratingData.ID, ratingData.MovieID, ratingData.RatingValue)

	return err
}
