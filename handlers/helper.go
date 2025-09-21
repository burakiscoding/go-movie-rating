package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type FileType uint8

const (
	Image FileType = iota
	Video
)

var acceptableFileCategories = map[uint8]FileType{
	0: Image, // Normal image
	1: Video, // Normal video
	2: Image, // Poster of the movie, main image on frontend
	3: Video, // Trailer of the movie, main video on frontend
}

var acceptableContentTypes = map[string]FileType{
	"image/jpeg": Image,
	"image/png":  Image,
	"image/webp": Image,
	"video/mp4":  Video,
	// "video/avi": Video,
	// "video/webm": Video,
}

// Checks the file's content type and category
// Returns true if content type and category is acceptable
func IsFileAcceptable(file multipart.File, category uint8) bool {
	// Read first 512 bytes to understand content type
	contentData := make([]byte, 512)
	if _, err := file.Read(contentData); err != nil {
		return false
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return false
	}

	// Check content type
	contentType := http.DetectContentType(contentData)
	ft1, ok := acceptableContentTypes[contentType]
	if !ok {
		return false
	}

	// Check category
	ft2, ok := acceptableFileCategories[category]
	if !ok {
		return false
	}

	// Matches file type and category
	// You can't upload an image as a trailer. Images can be poster or normal image.
	if ft1 != ft2 {
		return false
	}

	return true
}

func writeJSON(w http.ResponseWriter, code int, data any) {
	encoded, err := json.Marshal(data)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(encoded)
}

func WriteOK(w http.ResponseWriter, data any) {
	writeJSON(w, http.StatusOK, data)
}

func WriteError(w http.ResponseWriter, code int, message string) {
	writeJSON(w, code, map[string]string{"error": message})
}

func WriteForbidden(w http.ResponseWriter) {
	WriteError(w, http.StatusForbidden, "Forbidden action")
}

func WriteUnauthorized(w http.ResponseWriter) {
	WriteError(w, http.StatusUnauthorized, "Unauthenticated user")
}

func WriteNotFound(w http.ResponseWriter) {
	WriteError(w, http.StatusNotFound, "Requested resource not found")
}

func WriteBadRequest(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusBadRequest, err.Error())
}

func WriteServerError(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusInternalServerError, err.Error())
}

func WriteFailedValidation(w http.ResponseWriter, err error) {
	WriteError(w, http.StatusUnprocessableEntity, err.Error())
}

func WriteLargeRequestError(w http.ResponseWriter) {
	WriteError(w, http.StatusRequestEntityTooLarge, "Request entity too large")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashAndPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
