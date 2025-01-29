package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/service/seller"
	"online-shop-backend/pkg/response"
	"strconv"
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

	req := &seller.CreateSellerRequest{
		Name:  user.Name,
		Phone: user.Phone,
	}

	if _, err = h.services.Seller.CreateSeller(req); err != nil {
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
		err error
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &seller.GetSellerRequest{
		Id: intId,
	}

	resp, err := h.services.Seller.GetSeller(req)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) listSellers(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	req := &seller.ListSellersRequest{}

	resp, err := h.services.Seller.ListSellers(req)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (h *Handler) updateSeller(w http.ResponseWriter, r *http.Request) {
	var (
		err       error
		sellerEnt domain.Seller
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&sellerEnt); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.validator.Struct(sellerEnt); err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &seller.UpdateSellerRequest{
		Id:    intId,
		Name:  sellerEnt.Name,
		Phone: sellerEnt.Phone,
	}

	if _, err = h.services.Seller.UpdateSeller(req); err != nil {
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
	)

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	req := &seller.DeleteSellerRequest{
		Id: intId,
	}

	if _, err = h.services.Seller.DeleteSeller(req); err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(fmt.Sprintf("seller with ID %s was deleted", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
