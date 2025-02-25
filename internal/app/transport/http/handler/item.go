package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/service/item"
	"online-shop-backend/pkg/response"
	"strconv"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		req item.CreateItemRequest
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.services.Item.CreateItem(&req); err != nil {
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

	req := &item.GetItemRequest{
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

	req := &item.ListItemsRequest{}

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
		err error
		req item.UpdateItemRequest
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req.Id = intId

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.services.Item.UpdateItem(&req); err != nil {
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

	req := &item.DeleteItemRequest{
		Id: intId,
	}

	if _, err = h.services.Item.DeleteItem(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("item with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
