package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/service/buyer"
	"online-shop-backend/pkg/response"
	"strconv"
)

func (h *Handler) createBuyer(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		req buyer.CreateBuyerRequest
	)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(req); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err = h.services.Buyer.CreateBuyer(&req); err != nil {
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
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &buyer.GetBuyerRequest{
		Id: intId,
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

func (h *Handler) listBuyers(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	req := &buyer.ListBuyersRequest{}

	resp, err := h.services.Buyer.ListBuyers(req)
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
		err error
		req buyer.UpdateBuyerRequest
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

	if _, err = h.services.Buyer.UpdateBuyer(&req); err != nil {
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
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &buyer.DeleteBuyerRequest{
		Id: intId,
	}

	if _, err = h.services.Buyer.DeleteBuyer(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("buyer with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
