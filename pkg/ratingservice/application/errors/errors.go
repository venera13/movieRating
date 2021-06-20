package serviceerrors

import "errors"

var ErrorRequiredMovieID = errors.New("the movie id is required")
var ErrorRequiredRatingValue = errors.New("the rating value is required")
var ErrorMovieNotFound = errors.New("movie not found")
