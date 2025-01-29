package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/domain"
	buyer2 "online-shop-backend/internal/service/buyer"
	"online-shop-backend/pkg/response"
)

func (h *Handler) createBuyer(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		buyer domain.Buyer
	)

	if err = json.NewDecoder(r.Body).Decode(&buyer); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(buyer); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &buyer2.CreateBuyerRequest{
		Name:  buyer.Name,
		Phone: buyer.Phone,
	}

	if _, err = h.services.Buyer.CreateBuyer(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(fmt.Sprintf("buyer was created"))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) getBuyer(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		id  int
	)

	id, ok := r.Context().Value("id").(int)
	if !ok {
		response.NewErrorResponse(w, http.StatusBadRequest, "buyer id not found in context")
		return
	}

	req := &buyer2.GetBuyerRequest{
		Id: id,
	}

	resp, err := h.services.Buyer.GetBuyer(req)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) updateBuyer(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		id    int
		buyer domain.Buyer
	)

	id, ok := r.Context().Value("id").(int)
	if !ok {
		response.NewErrorResponse(w, http.StatusBadRequest, "buyer ID not found in context")
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&buyer); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(buyer); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &buyer2.UpdateBuyerRequest{
		Id:    id,
		Name:  "",
		Phone: "",
	}

	if _, err = h.services.Buyer.UpdateBuyer(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("buyer with ID %s was updated", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) deleteBuyer(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		id  int
	)

	id, ok := r.Context().Value("id").(int)
	if !ok {
		response.NewErrorResponse(w, http.StatusBadRequest, "buyer ID not found in context")
		return
	}

	req := &buyer2.DeleteBuyerRequest{Id: id}

	if _, err = h.services.Buyer.DeleteBuyer(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	if _, err := w.Write([]byte(fmt.Sprintf("buyer with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
