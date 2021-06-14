package provider

import (
	"ratingservice/pkg/ratingservice/application/data"
	"ratingservice/pkg/ratingservice/infrastructure"
)

type MovieProvider interface {
	Get(id string) (*data.Movie, error)
}

func NewMovieProvider() MovieProvider {
	return &infrastructure.MovieRepository{}
}
