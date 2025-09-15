package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/burakiscoding/go-movie-rating/db"
	"github.com/burakiscoding/go-movie-rating/handlers"
	"github.com/burakiscoding/go-movie-rating/middleware"
	"github.com/burakiscoding/go-movie-rating/stores"
)

func main() {
	db, err := db.NewSQL()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")

	movieStore := stores.NewMovieStore(db)
	movieHandler := handlers.NewMovieHandler(movieStore)
	userStore := stores.NewUserStore(db)
	userHandler := handlers.NewUserHandler(userStore)

	router := http.NewServeMux()
	router.HandleFunc("GET /movies", movieHandler.GetAll)
	router.HandleFunc("GET /movies/{id}", movieHandler.GetById)

	router.HandleFunc("POST /signUp", userHandler.SignUp)
	router.HandleFunc("POST /signIn", userHandler.SignIn)

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server listening on port 8080")
	server.ListenAndServe()
}
