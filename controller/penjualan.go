package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	client "sevima/db"
	db "sevima/db/sqlc"
	"sevima/utils"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func CreatePenjualan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		fmt.Println("Ini mah masuk")

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var jual db.CreatePenjualanParams

		var BeratSampah = r.FormValue("berat-sampah")
		var jenisSampah = r.FormValue("jenis-sampah")
		fmt.Print("Terisikah? ", BeratSampah)
		// if err := json.NewDecoder(r.Body).Decode(&jual); err != nil {
		// 	utils.ResponseJSON(w, err, http.StatusBadRequest)
		// 	return
		// }
		//var err error

		jual.BeratSampah, _ = strconv.Atoi(BeratSampah)
		jual.JenisSampah = jenisSampah
		switch jual.JenisSampah {
		case "buah":
			jual.HargaSampah = 500
			return
		case "sayur":
			jual.HargaSampah = 200
			return
		case "campuran":
			jual.HargaSampah = 200
			return
		default:
			fmt.Println("error")
		}

		jual.Keuntungan = int32(jual.BeratSampah) * jual.HargaSampah

		if err := client.DB.CreatePenjualan(ctx, jual); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}

func UpdateStatusPenjualan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var editStatus db.UpdateStatusPenjualanParams
		editStatus.Status = "success"

		if err := json.NewDecoder(r.Body).Decode(&editStatus); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.UpdateStatusPenjualan(ctx, editStatus); err != nil {
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
