package controller

import (
	"context"
	"fmt"
	"net/http"
	client "sevima/db"
	"sevima/utils"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func GetAllPembelian(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		pembelian, err := client.DB.GetAllRiwayatPembelian(ctx)

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, pembelian, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
}

func GetPembelianByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		ambilId := ps.ByName("id")
		ambilIdInt, _ := strconv.Atoi(ambilId)
		defer cancel()
		pembelian, err := client.DB.GetByIDRiwayatPembelian(ctx, int32(ambilIdInt))

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, pembelian, http.StatusOK)
		return
	}
	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
}

func DeletePembelian(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		idInt, _ := strconv.Atoi(id)
		if err := client.DB.DeleteRiwayatPembelian(ctx, int32(idInt)); err != nil {
			if strings.Contains(err.Error(), "foreign key constraint fails") {
				kesalahan := map[string]string{
					"error": "Data tidak dapat dihapus, masih ada data yang terikat",
				}

				utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
				return
			}
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
}
