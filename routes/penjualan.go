package routes

import (
	"sevima/controller"

	"github.com/julienschmidt/httprouter"
)

func PenjualanRouter(r *httprouter.Router) {
	r.POST("/penjualan/input", controller.CreatePenjualan)
	r.PUT("/penjualan/update", controller.UpdateStatusPenjualan)
}
