package infrastructure

import (
	"ratingservice/pkg/ratingservice/domain"
)

//func CreateRatingRepository(db *sql.DB) domain.RatingRepository {
//	return &DatabaseRepository{
//		db: db,
//	}
//}

//func CreateRepository(transaction Transaction) domain.RatingRepository {
//	return &DatabaseRepository{transaction: transaction}
//}

type DatabaseRepository struct {
	transaction Transaction
}

//type DatabaseRepository struct {
//	db *sql.DB
//}

func (ratingRepo *DatabaseRepository) AddRating(ratingData domain.Rating) error {
	query := "INSERT INTO rating(id, movie_id, rating_value) VALUES (?, ?, ?)"
	_, err := ratingRepo.transaction.Exec(query, ratingData.ID, ratingData.MovieID, ratingData.RatingValue)

	return err
}
