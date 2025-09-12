package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/burakiscoding/go-movie-rating/middleware"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /movies", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "List of movies")
	})
	router.HandleFunc("GET /movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		io.WriteString(w, "Movie "+id)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server listening on port 8080")
	server.ListenAndServe()
}
