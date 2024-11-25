package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/coupons", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createCoupon(w, r)
		case http.MethodGet:
			getCoupons(w, r)
		case http.MethodPatch:
			updateCoupons(w, r)
		case http.MethodDelete:
			deleteCoupon(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createProduct(w, r)
		case http.MethodGet:
			getProducts(w, r)
		}
	})

	http.HandleFunc("/applicable-coupons", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createCoupon(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/apply-coupon", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			applyCoupon(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

const (
	cartWise = iota
	productWise
	BxGy
)

type Coupon struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DiscountType int    `json:"discountType"`
	Discount     int    `json:"discount"`
	MinAmount    int    `json:"minAmount"`
	ProductX     int    `json:"buyX"`
	ProductY     int    `json:"getY"`
}

var (
	Coupons = make(map[int]Coupon)
	idCount = 1
	mutex   = &sync.Mutex{} // Protects access to items
)

var (
	Products  = make(map[int]Product)
	idProduct = 1
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductQuantity struct {
	Product  Product
	Quantity int
}

type FreeItems struct {
	Product  Product
	Quantity int
	FreeWith int
}

type Cart struct {
	Items     map[int]Product
	FreeItems map[int]Product
}
