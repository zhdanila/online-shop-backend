package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/domain"
	"online-shop-backend/pkg"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		item domain.Item
	)

	if err = json.NewDecoder(r.Body).Decode(&item); err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(item); err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Item.CreateItem(item); err != nil {
		pkg.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(fmt.Sprintf("item was created"))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) getItem(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		id      string
		product domain.Item
	)

	id, ok := r.Context().Value("id").(string)
	if !ok {
		pkg.NewErrorResponse(w, http.StatusBadRequest, "item id not found in context")
		return
	}

	if product, err = h.services.Item.GetItem(id); err != nil {
		pkg.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		id   string
		item domain.Item
	)

	id, ok := r.Context().Value("id").(string)
	if !ok {
		pkg.NewErrorResponse(w, http.StatusBadRequest, "item ID not found in context")
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&item); err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(item); err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Item.UpdateItem(id, item); err != nil {
		pkg.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
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
		id  string
	)

	id, ok := r.Context().Value("id").(string)
	if !ok {
		pkg.NewErrorResponse(w, http.StatusBadRequest, "item ID not found in context")
		return
	}

	if err = h.services.Item.DeleteItem(id); err != nil {
		pkg.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write([]byte(fmt.Sprintf("item with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
