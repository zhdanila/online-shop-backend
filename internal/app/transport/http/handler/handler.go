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

	mux.Handle("/seller", middleware.BasicAuth(secureMux))
	mux.Handle("/item", middleware.BasicAuth(secureMux))
	mux.Handle("/buyer", middleware.BasicAuth(secureMux))
	mux.Handle("/order", middleware.BasicAuth(secureMux))

	secureMux.Handle("POST /seller", http.HandlerFunc(h.createSeller))
	secureMux.Handle("GET /seller/{id}", http.HandlerFunc(h.getSeller))
	secureMux.Handle("PUT /seller/{id}", http.HandlerFunc(h.updateSeller))
	secureMux.Handle("DELETE /seller/{id}", http.HandlerFunc(h.deleteSeller))

	secureMux.Handle("POST /item", http.HandlerFunc(h.createItem))
	secureMux.Handle("GET /item/{id}", http.HandlerFunc(h.getItem))
	secureMux.Handle("PUT /item/{id}", http.HandlerFunc(h.updateItem))
	secureMux.Handle("DELETE /item/{id}", http.HandlerFunc(h.deleteItem))

	secureMux.Handle("POST /buyer", http.HandlerFunc(h.createBuyer))
	secureMux.Handle("GET /buyer/{id}", http.HandlerFunc(h.getBuyer))
	secureMux.Handle("PUT /buyer/{id}", http.HandlerFunc(h.updateBuyer))
	secureMux.Handle("DELETE /buyer/{id}", http.HandlerFunc(h.deleteBuyer))

	secureMux.Handle("POST /order", http.HandlerFunc(h.createOrder))
	secureMux.Handle("GET /order/{id}", http.HandlerFunc(h.getOrder))
	secureMux.Handle("GET /order", http.HandlerFunc(h.listOrders))
	secureMux.Handle("PUT /order/{id}", http.HandlerFunc(h.updateOrder))
	secureMux.Handle("DELETE /order/{id}", http.HandlerFunc(h.deleteOrder))

	return mux
}
