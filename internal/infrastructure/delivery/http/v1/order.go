package v1

import (
	"ddd-impl/internal/domains/customer"
	"ddd-impl/internal/domains/order"
	"ddd-impl/internal/domains/product"
	"ddd-impl/internal/infrastructure/repository/order"
	"encoding/json"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	ou order.Usecase
}

func NewOrderHandler(ou domains.Usecase) *OrderHandler {
	return &OrderHandler{
		ou: ou,
	}
}

// implement handler
func (o *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Products []*product.Product `json:"items"`
		customer customer.Customer  `json:"Customer"`
	}

	//  'urlxxxx?brand=apple&promo=true'
	// body

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// usecase order
	err := o.ou.Create(reqBody.Products, reqBody.customer)

	//// Membuat pesanan baru menggunakan OrderService
	//order, err := h.orderService.CreateOrder(reqBody.Products, reqBody.Customer)
	//if err != nil {
	//	http.Error(w, "Failed to create order", http.StatusInternalServerError)
	//	return
	//}

	// requst from fe

	orde := order.NewOrder(req)

	if err != nil {
		return
	}
}

func (o *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		status string `json:"status"`
	}

	rQuery := r.URL.Query().Get("brand")
	xrQuery := r.URL.Query().Get("promo")

	orderIDStr := r.URL.Query().Get("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}
	//  'urlxxxx?brand=apple&promo=true'
	// body

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// usecase order
	//err := o.ou.Create(reqBody.Products, reqBody.customer)
	order, err := o.ou.UpdateOrder(status, orderID, rQuery, xrQuery)
	//// Membuat pesanan baru menggunakan OrderService
	//order, err := h.orderService.CreateOrder(reqBody.Products, reqBody.Customer)
	//if err != nil {
	//	http.Error(w, "Failed to create order", http.StatusInternalServerError)
	//	return
	//}

	// requst from fe

	orde := order.NewOrder(req)

	if err != nil {
		return
	}
}
