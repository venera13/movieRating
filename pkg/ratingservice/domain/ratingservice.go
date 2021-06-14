package domain

type RatingRepository interface {
	Add(ratingData Rating) error
	Update(ratingData Rating) error
	Get(id string) (*Rating, error)
	GetRatingByMovieId(movieId string) (*Rating, error)
	Remove(ratingData Rating) error
}

type Rating struct {
	ID            string
	MovieID       string
	RatingValue   int64
	NumberRatings int64
	DeletedAt     string
}
