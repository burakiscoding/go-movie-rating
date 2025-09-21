package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/burakiscoding/go-movie-rating/db"
	"github.com/burakiscoding/go-movie-rating/handlers"
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
	router.HandleFunc("POST /movies/{id}/upload", movieHandler.UploadFile)
	router.HandleFunc("GET /movies/{id}/media/{name}", movieHandler.GetFile)
	router.HandleFunc("POST /movies/{id}/ratings", handlers.IsAuthenticated(movieHandler.AddRating))
	router.HandleFunc("GET /movies/{id}/ratings", movieHandler.GetRatings)

	router.HandleFunc("POST /user/signUp", userHandler.SignUp)
	router.HandleFunc("POST /user/signIn", userHandler.SignIn)
	router.HandleFunc("GET /user/profile", handlers.IsAuthenticated(userHandler.GetProfile))
	router.HandleFunc("PUT /user/profile", handlers.IsAuthenticated(userHandler.UpdateProfile))
	router.HandleFunc("GET /user/ratings", handlers.IsAuthenticated(userHandler.GetRatings))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port 8080")
	server.ListenAndServe()
}
