package controller

import (
	"context"
	"encoding/json"
	"net/http"
	client "sevima/db"
	db "sevima/db/sqlc"
	"sevima/utils"

	"github.com/julienschmidt/httprouter"
)

//POST
func RegisterUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var user db.RegisterUserParams

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.RegisterUser(ctx, user); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}
