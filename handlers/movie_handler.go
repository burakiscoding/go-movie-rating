package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/burakiscoding/go-movie-rating/stores"
	"github.com/burakiscoding/go-movie-rating/types"
	"github.com/go-playground/validator/v10"
)

type MovieHandler struct {
	store stores.MovieStore
}

func NewMovieHandler(store stores.MovieStore) MovieHandler {
	return MovieHandler{
		store: store,
	}
}

// GET /movies
func (h MovieHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := h.store.GetAll()
	if err != nil {
		WriteServerError(w, err)
		return
	}

	WriteOK(w, movies)
}

// GET /movies/{id}
func (h MovieHandler) GetById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		WriteBadRequest(w, err)
		return
	}

	movie, err := h.store.GetById(id)
	if err != nil {
		WriteNotFound(w)
		return
	}

	WriteOK(w, movie)
}

// POST /movies
func AddMovie(w http.ResponseWriter, r *http.Request) {}

// DELETE /movies/{id}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {}

// POST /movies/{id}/upload
func (h MovieHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// 1<<10 means 1KB, 1<<20 means 1MB, 1<<30 means 1GB
	// Set the max request body size to 500MB
	r.Body = http.MaxBytesReader(w, r.Body, 500<<20)

	// Limit memory: Creates temporary file in disk if the upload file is bigger than 100mb
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		WriteLargeRequestError(w)
		return
	}
	// Delete the temporary file
	defer r.MultipartForm.RemoveAll()

	// Get movie id
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		WriteBadRequest(w, err)
		return
	}

	// Check if the movie exists
	isMovieExists := h.store.IsMovieExists(id)
	if !isMovieExists {
		WriteNotFound(w)
		return
	}

	// Get files
	files, ok := r.MultipartForm.File["files"]
	if !ok {
		WriteNotFound(w)
		return
	}

	// Get types(poster, trailer, etc.)
	categories, ok := r.MultipartForm.Value["categories"]
	if !ok || len(categories) != len(files) {
		WriteNotFound(w)
		return
	}

	for i, header := range files {
		// Set the max file size to 100MB
		if header.Size > (100 << 20) {
			continue
		}

		file, err := header.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		category, err := strconv.Atoi(categories[i])
		if err != nil {
			return
		}

		// Checks the file type and category
		if !IsFileAcceptable(file, uint8(category)) {
			continue
		}

		fullPath := filepath.Join("uploads", idString, header.Filename)

		dst, err := os.Create(fullPath)
		if err != nil {
			continue
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			continue
		}

		// Save to db
		h.store.AddFile(id, header.Filename, uint8(category))

		// fmt.Fprintf(w, "Uploaded: %s\n", header.Filename)
	}

	WriteOK(w, map[string]string{"message": "Files uploaded"})
}

// GET movies/{id}/media/{name}
func (h MovieHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	movieId := r.PathValue("id")
	fileName := r.PathValue("name")

	// uploads/1/madame-web.jpg
	fullPath := filepath.Join("uploads", movieId, fileName)

	info, err := os.Stat(fullPath)
	if err != nil && os.IsNotExist(err) {
		WriteNotFound(w)
		return
	}

	if info.IsDir() {
		WriteNotFound(w)
		return
	}

	http.ServeFile(w, r, fullPath)
}

// POST /movies/{id}/ratings
func (h MovieHandler) AddRating(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(AuthUserId).(string)
	if !ok {
		return
	}

	movieIdString := r.PathValue("id")
	movieId, err := strconv.Atoi(movieIdString)
	if err != nil {
		WriteBadRequest(w, err)
		return
	}

	var body types.AddRatingBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteBadRequest(w, err)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		WriteFailedValidation(w, err)
		return
	}

	if err := h.store.AddRating(r.Context(), userId, movieId, body.Rating, body.Comment); err != nil {
		WriteServerError(w, err)
		return
	}

	WriteOK(w, map[string]string{"message": "Rating added"})
}

// GET /movies/{id}/ratings
func (h MovieHandler) GetRatings(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		WriteBadRequest(w, err)
		return
	}

	ratings, err := h.store.GetRatings(id)
	if err != nil {
		WriteServerError(w, err)
		return
	}

	WriteOK(w, ratings)
}
