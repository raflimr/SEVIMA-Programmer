package routes

import (
	"sevima/controller"

	"github.com/julienschmidt/httprouter"
)

func PembelianRouter(r *httprouter.Router) {
	r.POST("/pembelian/input", controller.CreatePembelian)
	r.PUT("/pembelian/update", controller.UpdateStatusPembelian)
}
