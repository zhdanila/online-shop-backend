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

func wrapMiddleware(h http.Handler) http.Handler {
	return middleware.BasicAuth(
		middleware.RecoverMiddleware(
			middleware.CORSMiddleware(h),
		),
	)
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("POST /seller", wrapMiddleware(http.HandlerFunc(h.createSeller)))
	mux.Handle("GET /seller", wrapMiddleware(http.HandlerFunc(h.listSellers)))
	mux.Handle("GET /seller/{id}", wrapMiddleware(http.HandlerFunc(h.getSeller)))
	mux.Handle("PUT /seller/{id}", wrapMiddleware(http.HandlerFunc(h.updateSeller)))
	mux.Handle("DELETE /seller/{id}", wrapMiddleware(http.HandlerFunc(h.deleteSeller)))

	mux.Handle("POST /item", wrapMiddleware(http.HandlerFunc(h.createItem)))
	mux.Handle("GET /item/{id}", wrapMiddleware(http.HandlerFunc(h.getItem)))
	mux.Handle("GET /item", wrapMiddleware(http.HandlerFunc(h.listItems)))
	mux.Handle("PUT /item/{id}", wrapMiddleware(http.HandlerFunc(h.updateItem)))
	mux.Handle("DELETE /item/{id}", wrapMiddleware(http.HandlerFunc(h.deleteItem)))

	mux.Handle("POST /buyer", wrapMiddleware(http.HandlerFunc(h.createBuyer)))
	mux.Handle("GET /buyer/{id}", wrapMiddleware(http.HandlerFunc(h.getBuyer)))
	mux.Handle("GET /buyer", wrapMiddleware(http.HandlerFunc(h.listBuyers)))
	mux.Handle("PUT /buyer/{id}", wrapMiddleware(http.HandlerFunc(h.updateBuyer)))
	mux.Handle("DELETE /buyer/{id}", wrapMiddleware(http.HandlerFunc(h.deleteBuyer)))

	mux.Handle("POST /order", wrapMiddleware(http.HandlerFunc(h.createOrder)))
	mux.Handle("GET /order/{id}", wrapMiddleware(http.HandlerFunc(h.getOrder)))
	mux.Handle("GET /order", wrapMiddleware(http.HandlerFunc(h.listOrders)))
	mux.Handle("PUT /order/{id}", wrapMiddleware(http.HandlerFunc(h.updateOrder)))
	mux.Handle("DELETE /order/{id}", wrapMiddleware(http.HandlerFunc(h.deleteOrder)))
	mux.Handle("POST /order/item", wrapMiddleware(http.HandlerFunc(h.addItemToOrder)))

	return mux
}
