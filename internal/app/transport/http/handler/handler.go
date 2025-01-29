package handler

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"online-shop-backend/internal/app/transport/http/middleware"
	"online-shop-backend/internal/service"
)

type Handler struct {
	services  *service.Service
	validator *validator.Validate
}

func NewHandler(services *service.Service, validate *validator.Validate) *Handler {
	return &Handler{
		services:  services,
		validator: validate,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	secureMux := http.NewServeMux()

	mux.Handle("/sellers", middleware.BasicAuth(secureMux))
	mux.Handle("/items", middleware.BasicAuth(secureMux))
	mux.Handle("/buyers", middleware.BasicAuth(secureMux))
	mux.Handle("/orders", middleware.BasicAuth(secureMux))

	secureMux.Handle("POST /sellers", http.HandlerFunc(h.createSeller))
	secureMux.Handle("GET /sellers/{id}", http.HandlerFunc(h.getSeller))
	secureMux.Handle("PUT /sellers/{id}", http.HandlerFunc(h.updateSeller))
	secureMux.Handle("DELETE /sellers/{id}", http.HandlerFunc(h.deleteSeller))

	secureMux.Handle("POST /items", http.HandlerFunc(h.createItem))
	secureMux.Handle("GET /items/{id}", http.HandlerFunc(h.getItem))
	secureMux.Handle("PUT /items/{id}", http.HandlerFunc(h.updateItem))
	secureMux.Handle("DELETE /items/{id}", http.HandlerFunc(h.deleteItem))

	secureMux.Handle("POST /buyers", http.HandlerFunc(h.createBuyer))
	secureMux.Handle("GET /buyers/{id}", http.HandlerFunc(h.getBuyer))
	secureMux.Handle("PUT /buyers/{id}", http.HandlerFunc(h.updateBuyer))
	secureMux.Handle("DELETE /buyers/{id}", http.HandlerFunc(h.deleteBuyer))

	secureMux.Handle("POST /orders", http.HandlerFunc(h.createOrder))
	secureMux.Handle("GET /orders/{id}", http.HandlerFunc(h.getOrder))
	secureMux.Handle("GET /orders", http.HandlerFunc(h.listOrders))
	secureMux.Handle("PUT /orders/{id}", http.HandlerFunc(h.updateOrder))
	secureMux.Handle("DELETE /orders/{id}", http.HandlerFunc(h.deleteOrder))

	return mux
}
