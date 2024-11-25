package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func createCoupon(w http.ResponseWriter, r *http.Request) {
	var newItem Coupon
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if newItem.DiscountType == 2 && (newItem.ProductX == 0 || newItem.ProductY == 0) && newItem.ProductX == newItem.ProductY {
		http.Error(w, "please specify product id to apply coupon", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	newItem.ID = idCount
	idCount++
	Coupons[newItem.ID] = newItem

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)

}

func getCoupons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		json.NewEncoder(w).Encode(Coupons)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if coupon, ok := Coupons[id]; ok {
		json.NewEncoder(w).Encode(coupon)
		return
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func updateCoupons(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedItem Coupon
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	updatedItem.ID = id

	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := Coupons[id]; ok {
		Coupons[id] = updatedItem
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Coupons[id])
		return

	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteCoupon(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	if coupon, ok := Coupons[id]; ok {
		delete(Coupons, id)
		json.NewEncoder(w).Encode(coupon)
		return
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}
