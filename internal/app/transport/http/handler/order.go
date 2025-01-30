package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/service/order"
	"online-shop-backend/pkg/response"
	"strconv"
)

func (h *Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		req order.CreateOrderRequest
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.services.Order.CreateOrder(&req); err != nil {
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
		err error
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &order.GetOrderRequest{
		Id: intId,
	}

	res, err := h.services.Order.GetOrder(req)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) listOrders(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	req := &order.ListOrderRequest{}

	res, err := h.services.Order.ListOrders(req)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) updateOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		req order.UpdateOrderRequest
	)

	id := r.PathValue("id")
	req.Id, err = strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.services.Order.UpdateOrder(&req); err != nil {
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
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &order.DeleteOrderRequest{
		Id: intId,
	}

	if _, err = h.services.Order.DeleteOrder(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("order with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) addItemToOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		req order.AddItemToOrderRequest
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err = h.services.Order.AddItemToOrder(&req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("order with item was added"))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
