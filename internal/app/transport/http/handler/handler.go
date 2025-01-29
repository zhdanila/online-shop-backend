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

	mux.Handle("POST /seller", middleware.BasicAuth(http.HandlerFunc(h.createSeller)))
	mux.Handle("GET /seller", middleware.BasicAuth(http.HandlerFunc(h.listSellers)))
	mux.Handle("GET /seller/{id}", middleware.BasicAuth(http.HandlerFunc(h.getSeller)))
	mux.Handle("PUT /seller/{id}", middleware.BasicAuth(http.HandlerFunc(h.updateSeller)))
	mux.Handle("DELETE /seller/{id}", middleware.BasicAuth(http.HandlerFunc(h.deleteSeller)))

	mux.Handle("POST /item", middleware.BasicAuth(http.HandlerFunc(h.createItem)))
	mux.Handle("GET /item/{id}", middleware.BasicAuth(http.HandlerFunc(h.getItem)))
	mux.Handle("GET /item", middleware.BasicAuth(http.HandlerFunc(h.listItems)))
	mux.Handle("PUT /item/{id}", middleware.BasicAuth(http.HandlerFunc(h.updateItem)))
	mux.Handle("DELETE /item/{id}", middleware.BasicAuth(http.HandlerFunc(h.deleteItem)))

	mux.Handle("POST /buyer", middleware.BasicAuth(http.HandlerFunc(h.createBuyer)))
	mux.Handle("GET /buyer/{id}", middleware.BasicAuth(http.HandlerFunc(h.getBuyer)))
	mux.Handle("GET /buyer", middleware.BasicAuth(http.HandlerFunc(h.listBuyers)))
	mux.Handle("PUT /buyer/{id}", middleware.BasicAuth(http.HandlerFunc(h.updateBuyer)))
	mux.Handle("DELETE /buyer/{id}", middleware.BasicAuth(http.HandlerFunc(h.deleteBuyer)))

	mux.Handle("POST /order", middleware.BasicAuth(http.HandlerFunc(h.createOrder)))
	mux.Handle("GET /order/{id}", middleware.BasicAuth(http.HandlerFunc(h.getOrder)))
	mux.Handle("GET /order", middleware.BasicAuth(http.HandlerFunc(h.listOrders)))
	mux.Handle("PUT /order/{id}", middleware.BasicAuth(http.HandlerFunc(h.updateOrder)))
	mux.Handle("DELETE /order/{id}", middleware.BasicAuth(http.HandlerFunc(h.deleteOrder)))
	mux.Handle("POST /order/item", middleware.BasicAuth(http.HandlerFunc(h.addItemToOrder)))

	return mux
}
