package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/go-market/app"
	"github.com/syrlramadhan/go-market/app/config"
	"github.com/syrlramadhan/go-market/app/model"
)

func main() {
	fmt.Println("Server running on http://localhost:9000")

	db := config.ConnectToDB()

	err := db.AutoMigrate(&model.MstUser{}, &model.MstAddresses{}, &model.MstProducts{})
	if err != nil {
        log.Fatal("Gagal melakukan migrasi:", err)
    }

	router := httprouter.New()
	handler := app.Routes(router, db)

	server := http.Server{
		Addr: ":9000",
		Handler: handler,
	}

	errServe := server.ListenAndServe()
	if errServe != nil {
		panic(errServe)
	}
}
