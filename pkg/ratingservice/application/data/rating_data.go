package data

type RateTheMovieInput struct {
	MovieId       string `json:"movie_id"`
	RatingValue   int64  `json:"rating_value"`
	NumberRatings int64  `json:"number_ratings"`
}
