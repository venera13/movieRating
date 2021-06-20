package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ratingservice/pkg/ratingservice/application/data"
)

type MovieRepository struct {
}

func (movieService *MovieRepository) Get(id string) (*data.Movie, error) {
	getMovieURL := fmt.Sprintf("http://localhost:8000/api/v1/movie/%s", id)
	resp, err := http.Get(getMovieURL)

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
