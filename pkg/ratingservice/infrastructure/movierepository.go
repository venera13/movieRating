package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ratingservice/pkg/ratingservice/domain"
)

type MovieRepository struct {
	transaction Transaction
}

func (movieAdapter *MovieRepository) Get(id string) (*domain.Movie, error) {
	getMovieUrl := fmt.Sprintf("http://localhost:8000/api/v1/movie/%s", id)
	resp, err := http.Get(getMovieUrl)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var movie domain.Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)

	if err != nil {
		return nil, err
	}

	return &movie, err
}
