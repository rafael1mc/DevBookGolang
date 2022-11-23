package controllers

import (
	"api/src/responses"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, struct {
		Server string `json:"server"`
	}{
		Server: "Running...",
	})
}
