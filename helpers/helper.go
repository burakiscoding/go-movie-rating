package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SayHello() {
	fmt.Println("Hi")
}

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
