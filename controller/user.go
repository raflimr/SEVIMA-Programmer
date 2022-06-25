package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	client "sevima/db"
	db "sevima/db/sqlc"
	"sevima/utils"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var user db.RegisterUserParams
		user.Role = "user"

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

func GetUsersByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		ambilId := ps.ByName("id")
		ambilIdInt, _ := strconv.Atoi(ambilId)
		defer cancel()
		var user db.User
		user, err := client.DB.GetUserByID(ctx, int32(ambilIdInt))

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, user, http.StatusOK)
		return
	}
	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
}

func GetUsersByUsernameAndPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		var users db.RegisterUserParams

		if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
			log.Println(err)
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		user, err := client.DB.GetUserByUsernameAndPassword(ctx, db.GetUserByUsernameAndPasswordParams{
			Username: users.Username,
			Password: users.Password,
		})

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, user, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
}

func UpdateSaldodanTotalSampah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var editSampahdanTotalSampah db.UpdateSaldodanTotalSampahParams

		if err := json.NewDecoder(r.Body).Decode(&editSampahdanTotalSampah); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.UpdateSaldodanTotalSampah(ctx, editSampahdanTotalSampah); err != nil {
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

func UpdateProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var editProfile db.UpdateProfileParams

		if err := json.NewDecoder(r.Body).Decode(&editProfile); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.UpdateProfile(ctx, editProfile); err != nil {
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
