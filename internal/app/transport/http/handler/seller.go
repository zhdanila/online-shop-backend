package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/domain"
	"online-shop-backend/pkg"
)

func (h *Handler) createSeller(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		user domain.Seller
	)

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(user); err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Seller.CreateSeller(user); err != nil {
		pkg.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
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
		id     string
		seller domain.Seller
	)

	id, ok := r.Context().Value("id").(string)
	if !ok {
		pkg.NewErrorResponse(w, http.StatusBadRequest, "user id not found in context")
		return
	}

	if seller, err = h.services.Seller.GetSeller(id); err != nil {
		pkg.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
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
		id   string
		user domain.Seller
	)

	id, ok := r.Context().Value("id").(string)
	if !ok {
		pkg.NewErrorResponse(w, http.StatusBadRequest, "seller ID not found in context")
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(user); err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.Seller.UpdateSeller(id, user); err != nil {
		pkg.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
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
		id  string
	)

	id, ok := r.Context().Value("id").(string)
	if !ok {
		pkg.NewErrorResponse(w, http.StatusBadRequest, "seller ID not found in context")
		return
	}

	if err = h.services.Seller.DeleteSeller(id); err != nil {
		pkg.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write([]byte(fmt.Sprintf("seller with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
