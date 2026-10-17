package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func Error(w http.ResponseWriter, status int, message, code string) {
	JSON(w, status, ErrorResponse{
		Error: message,
		Code:  code,
	})
}
func Success(w http.ResponseWriter, data any) {
	JSON(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusCreated, data)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func BadRequest(w http.ResponseWriter, message string) {
	Error(w, http.StatusBadRequest, message, "BAD_REQUEST")
}

func InternalError(w http.ResponseWriter) {
	Error(w, http.StatusInternalServerError, "Internal server error", "INTERNAL_ERROR")
}
