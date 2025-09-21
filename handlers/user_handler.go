package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/burakiscoding/go-movie-rating/stores"
	"github.com/burakiscoding/go-movie-rating/types"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	store stores.UserStore
}

func NewUserHandler(store stores.UserStore) UserHandler {
	return UserHandler{
		store: store,
	}
}

// POST /user/signUp
func (h UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var body types.AuthReqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteBadRequest(w, err)
		return
	}
	defer r.Body.Close()

	// Validate body
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		WriteFailedValidation(w, err)
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		WriteServerError(w, err)
		return
	}

	// Create user
	id, err := h.store.CreateUserAndProfile(r.Context(), body.Email, hashedPassword)
	if err != nil {
		WriteServerError(w, err)
		return
	}

	// Generate token
	token, err := CreateToken(id)
	if err != nil {
		WriteServerError(w, err)
		return
	}

	WriteOK(w, types.AuthResBody{Token: token, Email: body.Email})
}

// POST /user/signIn
func (h UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// Parse body
	var body types.AuthReqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		WriteBadRequest(w, err)
		return
	}
	defer r.Body.Close()

	// Validate Email and Password
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		WriteFailedValidation(w, err)
		return
	}

	// Get user
	user, err := h.store.GetUserByEmail(body.Email)
	if err != nil {
		WriteBadRequest(w, err)
		return
	}

	// Compare passwords
	if !CompareHashAndPassword(user.Password, body.Password) {
		WriteBadRequest(w, err)
		return
	}

	// Generate token
	token, err := CreateToken(user.Id)
	if err != nil {
		WriteServerError(w, err)
		return
	}

	WriteOK(w, types.AuthResBody{Token: token, Email: user.Email})
}

// PUT /user/profile
func (h UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var body types.UpdateProfileBody
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

	userId, ok := r.Context().Value(AuthUserId).(string)
	if !ok {
		WriteUnauthorized(w)
		return
	}

	if err := h.store.UpdateProfile(userId, body.FirstName, body.LastName, body.AboutMe); err != nil {
		WriteServerError(w, err)
		return
	}

	WriteOK(w, map[string]string{"message": "Updated profile"})
}

// GET /user/profile
func (h UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(AuthUserId).(string)
	if !ok {
		WriteUnauthorized(w)
		return
	}

	p, err := h.store.GetProfile(userId)
	if err != nil {
		WriteServerError(w, err)
		return
	}

	WriteOK(w, p)
}

// GET /user/ratings
func (h UserHandler) GetRatings(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(AuthUserId).(string)
	if !ok {
		WriteUnauthorized(w)
		return
	}

	ratings, err := h.store.GetRatings(userId)
	if err != nil {
		WriteServerError(w, err)
		return
	}

	WriteOK(w, ratings)
}
