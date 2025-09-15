package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/burakiscoding/go-movie-rating/helpers"
	"github.com/burakiscoding/go-movie-rating/stores"
	"github.com/burakiscoding/go-movie-rating/types"
	"github.com/go-playground/validator/v10"
)

// POST /sign-in
// POST /sign-up
// GET /user/profile
// PATCH /user/profile
// GET /user/ratings
type UserHandler struct {
	store stores.UserStore
}

func NewUserHandler(store stores.UserStore) UserHandler {
	return UserHandler{
		store: store,
	}
}

func (h UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var body types.AuthReqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: err.Error()})
		return
	}
	defer r.Body.Close()

	// Validate body
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: err.Error()})
		return
	}

	// Create user
	id, err := h.store.CreateUserAndProfile(r.Context(), body.Email, hashedPassword)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: err.Error()})
		return
	}

	// Generate token
	token, err := CreateToken(id)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: err.Error()})
		return
	}

	helpers.WriteJson(w, http.StatusOK, types.AuthResBody{Token: token, Email: body.Email})
}

func (h UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var body types.AuthReqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: err.Error()})
		return
	}
	defer r.Body.Close()

	// Validate Email and Password
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: err.Error()})
		return
	}

	// Get user
	user, err := h.store.GetUserByEmail(body.Email)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: err.Error()})
		return
	}

	// Compare passwords
	if !CompareHashAndPassword(user.Password, body.Password) {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: "Bad credentials"})
		return
	}

	// Generate token
	token, err := CreateToken(user.Id)
	if err != nil {
		helpers.WriteJson(w, http.StatusInternalServerError, types.ErrorResBody{Message: err.Error()})
	}

	helpers.WriteJson(w, http.StatusOK, types.AuthResBody{Token: token, Email: user.Email})
}
