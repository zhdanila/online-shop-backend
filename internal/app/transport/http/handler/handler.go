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
	mux.Handle("/products", middleware.BasicAuth(secureMux))
	mux.Handle("/buyers", middleware.BasicAuth(secureMux))
	mux.Handle("/orders", middleware.BasicAuth(secureMux))

	//secureMux.Handle("/sellers", http.HandlerFunc(h.createSeller))
	//secureMux.Handle("/sellers/{id}", http.HandlerFunc(h.getSeller))
	//secureMux.Handle("/sellers/{id}", http.HandlerFunc(h.updateSeller))
	//secureMux.Handle("/sellers/{id}", http.HandlerFunc(h.deleteSeller))
	//
	//secureMux.Handle("/products", http.HandlerFunc(h.createProduct))
	//secureMux.Handle("/products/{id}", http.HandlerFunc(h.getProduct))
	//secureMux.Handle("/products/{id}", http.HandlerFunc(h.updateProduct))
	//secureMux.Handle("/products/{id}", http.HandlerFunc(h.deleteProduct))
	//
	//secureMux.Handle("/buyers", http.HandlerFunc(h.createBuyer))
	//secureMux.Handle("/buyers/{id}", http.HandlerFunc(h.getBuyer))
	//secureMux.Handle("/buyers/{id}", http.HandlerFunc(h.updateBuyer))
	//secureMux.Handle("/buyers/{id}", http.HandlerFunc(h.deleteBuyer))
	//
	//secureMux.Handle("/orders", http.HandlerFunc(h.createOrder))
	//secureMux.Handle("/orders/{id}", http.HandlerFunc(h.getOrder))
	//secureMux.Handle("/orders", http.HandlerFunc(h.listOrders))
	//secureMux.Handle("/orders/{id}", http.HandlerFunc(h.updateOrder))
	//secureMux.Handle("/orders/{id}", http.HandlerFunc(h.deleteOrder))

	return mux
}
