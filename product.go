package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newItem Product
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	newItem.ID = idProduct
	idProduct++
	Products[newItem.ID] = newItem

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		json.NewEncoder(w).Encode(Products)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if product, ok := Products[id]; ok {
		json.NewEncoder(w).Encode(product)
		return
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}
