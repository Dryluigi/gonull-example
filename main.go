package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LukaGiorgadze/gonull"
	"github.com/gorilla/mux"
)

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Shop        *string `json:"shop"`
}

type UpdateProductReq struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Shop        *string  `json:"shop"`
}

type UpdateProductReqGoNull struct {
	Name        gonull.Nullable[string]  `json:"name"`
	Description gonull.Nullable[string]  `json:"description"`
	Price       gonull.Nullable[float64] `json:"price"`
	Shop        gonull.Nullable[string]  `json:"shop"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("Product name: %s\n", product.Name)
	fmt.Printf("Product description: %s\n", product.Description)
	fmt.Printf("Product price: %f\n", product.Price)
	fmt.Printf("Product price: %v\n", product.Shop)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product UpdateProductReqGoNull

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Pointer
	// if product.Name != nil {
	// 	fmt.Printf("Updated name to: %s \n", *product.Name)
	// }
	// if product.Description != nil {
	// 	fmt.Printf("Updated description to: %s \n", *product.Description)
	// }
	// if product.Price != nil {
	// 	fmt.Printf("Updated price to: %f \n", *product.Price)
	// }
	// if product.Shop != nil {
	// 	fmt.Printf("Updated shop to: %v \n", *product.Shop)
	// }

	// Gonull

	if product.Name.Present {
		fmt.Printf("Updated name to: %s \n", product.Name.Val)
	}
	if product.Description.Present {
		fmt.Printf("Updated description to: %s \n", product.Description.Val)
	}
	if product.Price.Present {
		fmt.Printf("Updated price to: %f \n", product.Price.Val)
	}
	if product.Shop.Present {
		if product.Shop.Valid {
			fmt.Printf("Updated shop to: %s \n", product.Shop.Val)
		} else {
			fmt.Printf("Updated shop to: null \n")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/products", CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/3", UpdateProduct).Methods("PATCH")

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
