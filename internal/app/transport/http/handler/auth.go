package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"online-shop-backend/internal/domain"
	"online-shop-backend/pkg"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validator.Struct(user); err != nil {
		pkg.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Auth.SignUp(user)
	if err != nil {
		pkg.NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte(fmt.Sprintf("person was created, id - %d", id))); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
