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

func GetAllPenjualan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		penjualan, err := client.DB.GetAllRiwayatPenjualan(ctx)

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, penjualan, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
}

func GetPenjualanByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		ambilId := ps.ByName("id")
		ambilIdInt, _ := strconv.Atoi(ambilId)
		defer cancel()
		penjualan, err := client.DB.GetByIDRiwayatPenjualan(ctx, int32(ambilIdInt))

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, penjualan, http.StatusOK)
		return
	}
	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
}

func DeletePenjualan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		idInt, _ := strconv.Atoi(id)
		if err := client.DB.DeleteRiwayatPenjualan(ctx, int32(idInt)); err != nil {
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
