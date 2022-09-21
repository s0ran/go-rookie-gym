package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/s0ran/go-rookie-gym/domain/user"
	"github.com/s0ran/go-rookie-gym/infrastructure"
)

type userRequest struct {
	Name string `json:"name"`
}

type userResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var req userRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db, err := infrastructure.NewDB()
	if err != nil {
		fmt.Printf("failed to connect db: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repo := infrastructure.NewRepository(db)
	id, err := repo.PostUser(r.Context(), user.NewUser(req.Name))
	if err != nil {
		fmt.Printf("failed to post user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res := userResponse{
		ID:   id,
		Name: req.Name,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
