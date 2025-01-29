package handler

import (
	"github.com/go-playground/validator/v10"
	"net/http"
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

	mux.HandleFunc("POST /sign-up", h.signUp)

	return mux
}
