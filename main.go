package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price int `json:"price"`
	Stock int    `json:"stock"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

var product = []Product{
	{ID: 1, Name: "Laptop", Price: 10000000, Stock: 10},
	{ID: 2, Name: "Smartphone", Price: 5000000, Stock: 25},
	{ID: 3, Name: "Tablet", Price: 7500000, Stock: 15},
}

var category = []Category{
	{ID: 1, Name: "Electronics", Description: "Devices and gadgets"},
	{ID: 2, Name: "Home Appliances", Description: "Appliances for home use"},
	{ID: 3, Name: "Books", Description: "Various kinds of books"},
}


func main(){

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") 
		json.NewEncoder(w).Encode(map[string]string{"status": "OK", "message": "Api Running"})
	})

	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
		} else if r.Method == http.MethodPost {
			var newProduct Product
			err := json.NewDecoder(r.Body).Decode(&newProduct)

			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			newProduct.ID = len(product) + 1
			product = append(product, newProduct)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newProduct)
		}
	 
	})

	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			updateProduct(w, r)
		} else if r.Method == http.MethodDelete {
			deleteProduct(w, r)
		} else if r.Method == http.MethodGet {
			getProductByID(w, r)
		}
	})

// * **GET** `/categories` → Ambil semua kategori
// * **POST** `/categories` → Tambah kategori
// * **PUT** `/categories/{id}` → Update kategori
// * **GET** `/categories/{id}` → Ambil detail satu kategori
// * **DELETE** `/categories/{id}` → Hapus kategori

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
		} else if r.Method == http.MethodPost {
			var newCategory Category

			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			newCategory.ID = len(category)+1
			category = append(category, newCategory)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCategory)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			updateCategory(w, r)
		} else if r.Method == http.MethodDelete {
			deleteCategory(w, r)
		} else if r.Method == http.MethodGet {
			getCategoryByID(w, r)
		}
	})



	fmt.Println("Server running in localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, p := range product {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	
	http.Error(w, "Product Not Found", http.StatusNotFound)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updateProduct Product

	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i, p := range product {
		if p.ID == id {
			if updateProduct.Name != "" {
				product[i].Name = updateProduct.Name
			}
			if updateProduct.Price != 0 {
				product[i].Price = updateProduct.Price
			}
			if updateProduct.Stock != 0 {
				product[i].Stock = updateProduct.Stock
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product[i])
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for i, p := range product {
		if p.ID == id {
			product = append(product[:i], product[i+1:]...)
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Product Deleted"})
			return
		}
	}

	http.Error(w, "Product Not Found", http.StatusNotFound)
}



// Category Handlers
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, p := range category {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	
	http.Error(w, "Product Not Found", http.StatusNotFound)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i, p := range category {
		if p.ID == id {
			if updateCategory.Name != "" {
				category[i].Name = updateCategory.Name
			}

			if updateCategory.Description != "" {
				category[i].Description = updateCategory.Description
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category[i])
			return
		}
	}

	http.Error(w, "Category Not Found", http.StatusNotFound)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, p := range category {
		if p.ID == id {
			category = append(category[:i], category[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Category Deleted"})
			return
		}
	}

	http.Error(w, "Category Not Found", http.StatusNotFound)
}