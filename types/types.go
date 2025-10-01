package types

import (
	"encoding/json"
	"time"
)

type Movie struct {
	Id                int           `json:"id"`
	Name              string        `json:"name"`
	Poster            string        `json:"poster"`
	Rating            float64       `json:"rating"`
	ReleaseDate       time.Time     `json:"releaseDate"`
	DurationInMinutes time.Duration `json:"durationInMinutes"`
}

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Actor struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type MovieDetail struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Rating            float64         `json:"rating"`
	ReleaseDate       time.Time       `json:"releaseDate"`
	DurationInMinutes time.Duration   `json:"durationInMinutes"`
	Genres            json.RawMessage `json:"genres"`
	Actors            json.RawMessage `json:"actors"`
	Medias            json.RawMessage `json:"medias"`
}

type MovieDetailResBody struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`
	Rating            float64         `json:"rating"`
	ReleaseDate       time.Time       `json:"releaseDate"`
	DurationInMinutes time.Duration   `json:"durationInMinutes"`
	Genres            json.RawMessage `json:"genres"`
	Actors            json.RawMessage `json:"actors"`
	Medias            []Media         `json:"medias"`
}

type Media struct {
	Name     string `json:"name"`
	Category uint8  `json:"type"`
}

type AuthReqBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

type AuthResBody struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

type TokenPayload struct {
	Id string `json:"id"`
}

type User struct {
	Id       string
	Email    string
	Password string
}

type MovieRating struct {
	Id        int     `json:"id"`
	MovieId   int     `json:"movieId"`
	Rating    float64 `json:"rating"`
	Comment   string  `json:"comment"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
}

type AddRatingBody struct {
	Rating  float64 `json:"rating" validate:"required,min=1,max=10"`
	Comment string  `json:"comment" validate:"required,max=255"`
}

type UpdateProfileBody struct {
	FirstName string `json:"firstName" validate:"required,max=30"`
	LastName  string `json:"lastName" validate:"required,max=30"`
	AboutMe   string `json:"aboutMe" validate:"required,max=255"`
}

type MovieRatingOfUser struct {
	Id        int     `json:"id"`
	MovieId   int     `json:"movieId"`
	MovieName string  `json:"movieName"`
	Rating    float64 `json:"rating"`
	Comment   string  `json:"comment"`
}

type Profile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	AboutMe   string `json:"aboutMe"`
}
