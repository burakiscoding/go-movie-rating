package handlers

import (
	"net/http"
	"strconv"

	"github.com/burakiscoding/go-movie-rating/helpers"
	"github.com/burakiscoding/go-movie-rating/stores"
)

// GET /movies
// GET /movies/{id}
// POST /movies
// GET /movies/{id}/ratings
// POST /movies/{id}/ratings
// POST /genres
// POST /actors
type MovieHandler struct {
	store stores.MovieStore
}

func NewMovieHandler(store stores.MovieStore) MovieHandler {
	return MovieHandler{
		store: store,
	}
}

func (h MovieHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := h.store.GetAll()
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, movies)
}

func (h MovieHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	movie, err := h.store.GetById(id)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJson(w, http.StatusOK, movie)
}
