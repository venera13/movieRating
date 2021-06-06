package repository

import (
	"database/sql"
	"ratingservice/pkg/ratingservice/domain"
)

func CreateRatingRepository(db *sql.DB) domain.RatingRepository {
	return &DatabaseRepository{
		db: db,
	}
}

type DatabaseRepository struct {
	db *sql.DB
}

func (ratingRepo *DatabaseRepository) AddRating(ratingData domain.Rating) error {
	query := "INSERT INTO rating(id, movie_id, rating_value) VALUES (?, ?, ?)"
	_, err := ratingRepo.db.Exec(query, ratingData.ID, ratingData.MovieID, ratingData.RatingValue)

	return err
}
