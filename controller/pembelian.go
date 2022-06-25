package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	client "sevima/db"
	db "sevima/db/sqlc"
	"sevima/utils"

	"github.com/julienschmidt/httprouter"
)

func CreatePembelian(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var beli db.CreatePembelianParams

		if err := json.NewDecoder(r.Body).Decode(&beli); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.CreatePembelian(ctx, beli); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}

func UpdateStatusPembelian(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var editStatus db.UpdateStatusPembelianParams
		editStatus.Status = "success"

		if err := json.NewDecoder(r.Body).Decode(&editStatus); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.UpdateStatusPembelian(ctx, editStatus); err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
}
