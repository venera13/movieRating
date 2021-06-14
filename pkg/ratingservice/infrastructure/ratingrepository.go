package infrastructure

import (
	"database/sql"
	"errors"
	"ratingservice/pkg/ratingservice/domain"
)

type DatabaseRepository struct {
	transaction Transaction
}

func (ratingRepo *DatabaseRepository) Add(ratingData domain.Rating) error {
	query := "INSERT INTO rating(id, movie_id, rating_value, number_ratings) VALUES (?, ?, ?, ?)"
	_, err := ratingRepo.transaction.Exec(query, ratingData.ID, ratingData.MovieID, ratingData.RatingValue, ratingData.NumberRatings)

	return err
}

func (ratingRepo *DatabaseRepository) Update(ratingData domain.Rating) error {
	query := "UPDATE rating SET rating_value = ?, number_ratings = ? WHERE id = ?"
	_, err := ratingRepo.transaction.Exec(query, ratingData.RatingValue, ratingData.NumberRatings, ratingData.ID)

	return err
}

func (ratingRepo *DatabaseRepository) Get(id string) (*domain.Rating, error) {
	var rating domain.Rating
	rating.ID = id
	query := "SELECT movie_id, rating_value FROM rating WHERE id = ? "
	row := ratingRepo.transaction.QueryRow(query, id)

	err := row.Scan(&rating)

	if errors.Is(err, sql.ErrNoRows) {
		return &rating, domain.ErrorRatingNotFound
	}

	return &rating, nil
}

func (ratingRepo *DatabaseRepository) GetRatingByMovieId(movieId string) (*domain.Rating, error) {
	var rating domain.Rating
	rating.MovieID = movieId
	query := "SELECT id, rating_value, number_ratings FROM rating WHERE movie_id = ? "
	row := ratingRepo.transaction.QueryRow(query, movieId)

	err := row.Scan(&rating.ID, &rating.RatingValue, &rating.NumberRatings)

	if errors.Is(err, sql.ErrNoRows) {
		return &rating, domain.ErrorMovieNotFound
	}

	return &rating, nil
}

func (ratingRepo *DatabaseRepository) Remove(ratingData domain.Rating) error {
	query := "UPDATE rating SET deleted_at = ? WHERE id = ?"
	_, err := ratingRepo.transaction.Exec(query, ratingData.DeletedAt, ratingData.ID)

	return err
}
