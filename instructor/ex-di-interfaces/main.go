package main

import (
	"encoding/json"
	"ex-di-interfaces/repository"
	"ex-di-interfaces/service"
	"log"
	"net/http"
)

type ProductRouter struct {
	service service.DefaultProductService
}

func NewProductRouter(ps service.DefaultProductService) ProductRouter {
	return ProductRouter{ps}
}

func main() {
	// creating instances of dependencies
	dbRepo := repository.NewDbProductRepo()
	// wiring your application (configuration code)
	ps := service.NewProductService(dbRepo)
	router := NewProductRouter(ps)

	http.HandleFunc("/products", router.productsHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (pr ProductRouter) productsHandler(w http.ResponseWriter, r *http.Request) {
	products := pr.service.GetAllProducts()
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Some error occoured"))
	}
}
