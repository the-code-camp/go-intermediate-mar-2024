package main

import (
	"encoding/json"
	"ex-di-interfaces/model"
	"ex-di-interfaces/repository"
	"ex-di-interfaces/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Employee struct {
	Id     int     `json:"id"`
	Name   string  `json:"employee_name"`
	Salary float32 `json:"-"`
}

//{"Id":100,"Name":"emp-1","Salary":1234.56}
//{"id":100,"employee_name":"emp-1","salary":1234.56}

func main() {

	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/products", productsHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	dbRepo := repository.NewDbProductRepo()
	ps := service.NewProductService(dbRepo)
	products := ps.GetAllProducts()
	// w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(products); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Some error occoured"))
	}
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	w.Header().Add("Content-Type", "application/json")

	empId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte(err.Error()))
		// errorResponse := map[string]string{
		// 	"mesage": err.Error(),
		// }
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
	} else {
		emp := Employee{empId, "emp-1", 1234.56}
		if err := json.NewEncoder(w).Encode(emp); err != nil {
			// send a bad response code
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Some error occoured"))
		}
	}
}

func Start() {
	// creating instances of dependencies
	dbRepo := repository.NewDbProductRepo()
	// inMemRepo := repository.NewInMemoryProductRepo()

	// wiring your application (configuration code)
	ps := service.NewProductService(dbRepo)

	// Printing all products
	printProducts(ps.GetAllProducts())

	newlyAddedProduct := ps.AddProduct("new product", model.BOOKS, 109.99)

	// printing the newly added product
	printProduct(newlyAddedProduct.Id, ps)

}

func printProducts(products []model.Product) {
	fmt.Println("### Printing Products ###")
	for _, p := range products {
		fmt.Println(p)
	}
	fmt.Println("### End ###")
}

func printProduct(id int, ps service.DefaultProductService) {
	if product := ps.GetProductById(id); product != nil {
		fmt.Println("PRODUCT: ", product)
	} else {
		fmt.Println(fmt.Sprintf("Product with id %d does not exist", id))
	}
}
