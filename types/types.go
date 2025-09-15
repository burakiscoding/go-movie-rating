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
	Poster            string          `json:"poster"`
	Rating            float64         `json:"rating"`
	ReleaseDate       time.Time       `json:"releaseDate"`
	DurationInMinutes time.Duration   `json:"durationInMinutes"`
	Genres            json.RawMessage `json:"genres"`
	Actors            json.RawMessage `json:"actors"`
}

type AuthReqBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

type AuthResBody struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

type ErrorResBody struct {
	Message string `json:"message"`
}

type TokenPayload struct {
	Id string `json:"id"`
}

type User struct {
	Id       string
	Email    string
	Password string
}
