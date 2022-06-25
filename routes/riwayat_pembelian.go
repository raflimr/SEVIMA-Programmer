package routes

import (
	"sevima/controller"

	"github.com/julienschmidt/httprouter"
)

func RiwayatPembelianRouter(r *httprouter.Router) {
	r.GET("/riwayat-pembelian", controller.GetAllPembelian)
	r.GET("/riwayat-pembelian/:id", controller.GetPembelianByID)
	r.DELETE("/riwayat-pembelian", controller.DeletePembelian)
}
