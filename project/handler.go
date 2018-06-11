package project

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// InitRouters init project routers
func InitRouters(router *httprouter.Router, service *Service) {
	router.GET("/v1/projects", index(service))
}

func index(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ps, err := service.FindAll()

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(ps)
	}
}
