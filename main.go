package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Category    *Category `json:"category"`
}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var products []Product

// GetProducts /**
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		return
	}
}

/**
 * @api {get} /products Get all products
 */
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range products {
		number, _ := strconv.ParseUint(params["id"], 10, 32)
		if item.Id == int(number) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

/**
 * @api {delete} /products/{id} Delete product
 */
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		number, _ := strconv.ParseUint(params["id"], 10, 32)
		if item.Id == int(number) {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}

/**
 * @api {post} /products Create product
 */
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.Id = len(products) + 1
	products = append(products, product)
	json.NewEncoder(w).Encode(products)
}

/**
 * @api {put} /products/{id} Update product
 */
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		id, _ := strconv.ParseUint(params["id"], 10, 32)
		if item.Id == int(id) {
			products = append(products[:index], products[index+1:]...)
			var product Product
			product.Id = int(id)
			_ = json.NewDecoder(r.Body).Decode(&product)
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			break
		}
	}

}

func main() {
	r := mux.NewRouter()

	products = append(products, Product{Id: 1, Name: "Product 1", Description: "Product 1 description", Price: 100, Category: &Category{Id: 1, Name: "Category 1"}})
	products = append(products, Product{Id: 2, Name: "Product 2", Description: "Product 2 description", Price: 200, Category: &Category{Id: 2, Name: "Category 2"}})
	products = append(products, Product{Id: 3, Name: "Product 3", Description: "Product 3 description", Price: 300, Category: &Category{Id: 3, Name: "Category 3"}})

	r.HandleFunc("/products", GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")

	fmt.Printf("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
