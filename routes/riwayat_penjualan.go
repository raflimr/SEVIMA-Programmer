package routes

import (
	"sevima/controller"

	"github.com/julienschmidt/httprouter"
)

func RiwayatPenjualanRouter(r *httprouter.Router) {
	r.GET("/riwayat-penjualan", controller.GetAllPenjualan)
	r.GET("/riwayat-penjualan/:id", controller.GetPenjualanByID)
	r.DELETE("/riwayat-penjualan", controller.DeletePenjualan)
}
