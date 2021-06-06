package adapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	domain "ratingservice/pkg/ratingservice/domain/adapter/movieservice"
)

func CreateMovieAdapter() domain.MovieAdapter {
	return &Adapter{}
}

type Adapter struct {
}

func (movieAdapter *Adapter) Get(id string) (*domain.Movie, error) {
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
