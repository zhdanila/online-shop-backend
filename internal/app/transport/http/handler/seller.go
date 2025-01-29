package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/domain"
	"online-shop-backend/pkg/response"
)

func (h *Handler) createSeller(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		user domain.Seller
	)

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(user); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Seller.CreateSeller(user); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(fmt.Sprintf("seller was created"))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) getSeller(w http.ResponseWriter, r *http.Request) {
	var (
		err    error
		id     int
		seller domain.Seller
	)

	id, ok := r.Context().Value("id").(int)
	if !ok {
		response.NewErrorResponse(w, http.StatusBadRequest, "seller ID not found in context")
		return
	}

	if seller, err = h.services.Seller.GetSeller(id); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(seller); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) updateSeller(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		id   int
		user domain.Seller
	)

	id, ok := r.Context().Value("id").(int)
	if !ok {
		response.NewErrorResponse(w, http.StatusBadRequest, "seller ID not found in context")
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(user); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Seller.UpdateSeller(id, user); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("seller with ID %s was updated", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) deleteSeller(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		id  int
	)

	id, ok := r.Context().Value("id").(int)
	if !ok {
		response.NewErrorResponse(w, http.StatusBadRequest, "seller ID not found in context")
		return
	}

	if err = h.services.Seller.DeleteSeller(id); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("seller with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
