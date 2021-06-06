package transport

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	service "ratingservice/pkg/ratingservice/application"
	"ratingservice/pkg/ratingservice/application/data"
	"time"
)

type Server struct {
	ratingService service.RatingService
}

func NewServer(service service.RatingService) *Server {
	return &Server{
		ratingService: service,
	}
}

func Router(srv *Server) http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/rate-movie", srv.rateTheMovie).Methods(http.MethodPost)

	return logMiddleware(r)
}

func (srv *Server) rateTheMovie(w http.ResponseWriter, r *http.Request) {
	requestData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		processError(w, err)

		return
	}
	defer r.Body.Close()

	var rateTheMovieInput data.RateTheMovieInput
	err = json.Unmarshal(requestData, &rateTheMovieInput)

	if err != nil {
		processError(w, err)

		return
	}

	err = srv.ratingService.RateTheMovie(&rateTheMovieInput)
	if err != nil {
		processError(w, err)

		return
	}
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
			"time":       time.Since(start),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}

func processError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
