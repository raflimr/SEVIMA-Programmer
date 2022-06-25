package routes

import (
	"sevima/controller"

	"github.com/julienschmidt/httprouter"
)

func UserRouter(r *httprouter.Router) {
	r.POST("/user/register", controller.RegisterUser)
	r.GET("/user/:id", controller.GetUsersByID)
	r.POST("/user/login", controller.GetUsersByUsernameAndPassword)
	r.PUT("/user/update-admin", controller.UpdateSaldodanTotalSampah)
	r.PUT("/user/profile", controller.UpdateProfile)
}
