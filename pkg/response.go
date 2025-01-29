package pkg

import "net/http"

func NewErrorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	w.WriteHeader(statusCode)
	if _, err := w.Write([]byte(errorString)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
