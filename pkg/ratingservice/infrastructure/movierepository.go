package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ratingservice/pkg/ratingservice/application/data"
)

type MovieRepository struct {
	transaction Transaction
}

func (movieService *MovieRepository) Get(id string) (*data.Movie, error) {
	getMovieUrl := fmt.Sprintf("http://localhost:8000/api/v1/movie/%s", id)
	resp, err := http.Get(getMovieUrl)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var movie data.Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)

	if err != nil {
		return nil, err
	}

	return &movie, err
}
