package domain

type RatingRepository interface {
	AddRating(ratingData Rating) error
}

type Rating struct {
	ID          string
	MovieID     string
	RatingValue string
}
