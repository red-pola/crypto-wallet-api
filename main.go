package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/red-pola/crypto-wallet-api/project"
)

func main() {
	router := httprouter.New()

	// Projects
	projectRepository := project.NewInMemRepository()
	projectService := project.NewService(projectRepository)
	project.InitRouters(router, projectService)

	// Ping
	router.GET("/status", func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		res.WriteHeader(http.StatusOK)
	})

	log.Println("CryptoWalletAPI is listening on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
