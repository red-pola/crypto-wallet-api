package project

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/red-pola/crypto-wallet-api/entity"
)

// InitRouters init project routers
func InitRouters(router *httprouter.Router, service *Service) {
	router.GET("/api/v1/projects", index(service))
	router.POST("/api/v1/projects", save(service))
	router.PUT("/api/v1/projects/:id", update(service))
	router.DELETE("/api/v1/projects/:id", remove(service))
}

func index(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		ps, err := service.FindAll()

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(res).Encode(ps); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func save(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		var p *entity.Project

		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := service.Save(p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(res).Encode(p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func update(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		var p *entity.Project
		id := entity.StringToID(params.ByName("id"))

		if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := service.Update(entity.ID(id), p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(res).Encode(p); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func remove(service *Service) httprouter.Handle {
	return func(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
		id := entity.StringToID(params.ByName("id"))

		if err := service.Delete(id); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusOK)
	}
}
