package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/service/order"
	"online-shop-backend/pkg/response"
	"strconv"
)

func (h *Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		orderEnt domain.Order
	)

	if err = json.NewDecoder(r.Body).Decode(&orderEnt); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(orderEnt); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &order.CreateOrderRequest{
		BuyerID:    orderEnt.BuyerID,
		TotalPrice: orderEnt.TotalPrice,
	}

	if _, err = h.services.Order.CreateOrder(req); err != nil {
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
		err      error
		orderEnt domain.Order
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&orderEnt); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(orderEnt); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &order.UpdateOrderRequest{
		Id:         intId,
		BuyerID:    orderEnt.BuyerID,
		TotalPrice: orderEnt.TotalPrice,
	}

	if _, err = h.services.Order.UpdateOrder(req); err != nil {
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

	req := &order.AddItemToOrderRequest{
		OrderID:  orderItem.OrderID,
		ItemID:   orderItem.ItemID,
		Quantity: orderItem.Quantity,
		Price:    orderItem.Price,
	}

	if _, err = h.services.Order.AddItemToOrder(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write([]byte(fmt.Sprintf("order with item was added"))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
