package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/domain"
	"online-shop-backend/pkg/response"
)

func (h *Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		order domain.Order
	)

	if err = json.NewDecoder(r.Body).Decode(&order); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(order); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Order.CreateOrder(order); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(fmt.Sprintf("order was created"))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) getOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		id    int
		order domain.Order
	)

	id, ok := r.Context().Value("id").(int)
	if !ok {
		response.NewErrorResponse(w, http.StatusBadRequest, "order id not found in context")
		return
	}

	if order, err = h.services.Order.GetOrder(id); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) listOrders(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		orders []domain.Order
	)

	if orders, err = h.services.Order.ListOrders(); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) updateOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		id    int
		order domain.Order
	)

	id, ok := r.Context().Value("id").(int)
	if !ok {
		response.NewErrorResponse(w, http.StatusBadRequest, "order ID not found in context")
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&order); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(order); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Order.UpdateOrder(id, order); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("order with ID %s was updated", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) deleteOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		id  int
	)

	id, ok := r.Context().Value("id").(int)
	if !ok {
		response.NewErrorResponse(w, http.StatusBadRequest, "order ID not found in context")
		return
	}

	if err = h.services.Order.DeleteOrder(id); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write([]byte(fmt.Sprintf("order with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) addItemToOrder(w http.ResponseWriter, r *http.Request) {
	var orderItem domain.OrderItems

	err := json.NewDecoder(r.Body).Decode(&orderItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.services.Order.AddItemToOrder(orderItem); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write([]byte(fmt.Sprintf("order with item was added"))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
