package handler

import (
	"encoding/json"
	"net/http"
)

type userRequest struct {
	Name string `json:"name"`
}

type userResponse struct {
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var req userRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse{Name: req.Name})
}
