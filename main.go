package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sevima/config"
	"sevima/controller"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Connect To DB")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := httprouter.New()
	router.POST("/user/register", controller.RegisterUser)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
