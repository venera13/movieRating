package serviceerrors

import "errors"

var RequiredMovieIdError = errors.New("the movie id is required")
var RequiredRatingValueError = errors.New("the rating value is required")
var MovieNotFound = errors.New("movie not found")
