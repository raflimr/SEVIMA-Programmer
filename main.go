package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sevima/config"
	client "sevima/db"
	sqlClient "sevima/db/sqlc"
	"sevima/routes"

	"github.com/julienschmidt/httprouter"
)

func CekAdmin() {
	admin, err := client.DB.GetAdminUser(context.TODO())
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if admin.ID == 0 {
		client.DB.RegisterUser(context.TODO(), sqlClient.RegisterUserParams{
			Username: "admin",
			Password: "admin",
			Role:     "admin",
		})
	}
}

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
	client.DB = sqlClient.New(db)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := httprouter.New()

	routes.UserRouter(router)
	routes.RiwayatPenjualanRouter(router)
	routes.RiwayatPembelianRouter(router)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
