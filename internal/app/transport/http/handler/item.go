package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/domain"
	item2 "online-shop-backend/internal/service/item"
	"online-shop-backend/pkg/response"
	"strconv"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		item domain.Item
	)

	if err = json.NewDecoder(r.Body).Decode(&item); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(item); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &item2.CreateItemRequest{
		SellerID:    item.SellerID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
	}

	if _, err = h.services.Item.CreateItem(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(fmt.Sprintf("item was created"))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) getItem(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &item2.GetItemRequest{
		Id: intId,
	}

	resp, err := h.services.Item.GetItem(req)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) listItems(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	req := &item2.ListItemsRequest{}

	resp, err := h.services.Item.ListItems(req)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		item domain.Item
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&item); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(item); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &item2.UpdateItemRequest{
		Id:          intId,
		SellerID:    item.SellerID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
	}

	if _, err = h.services.Item.UpdateItem(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("item with ID %s was updated", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &item2.DeleteItemRequest{
		Id: intId,
	}

	if _, err = h.services.Item.DeleteItem(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write([]byte(fmt.Sprintf("item with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
