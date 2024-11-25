package main

import (
	"net/http"
)

func applyCoupon(w http.ResponseWriter, r *http.Request) {
	//read coupon code and apply coupon code to the products in cart,
	//if coupon is eligible to product discount will apply
}

func applicableCoupons(w http.ResponseWriter, r *http.Request) {
	//iterate through all coupons and returns list of coupons which are applicable
}
