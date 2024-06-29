package utils

import (
	"encoding/json"
	"net/http"
	"technical-challenge/internal/core/domain/constants"
	"technical-challenge/internal/core/domain/models"
)

func Response(w http.ResponseWriter, devResponse models.DevResponse) {
	body, err := json.Marshal(devResponse.Response)
	if err != nil {
		body, _ := json.Marshal(models.ResponseWithResult{
			Response: &models.Response{
				Message: constants.REQUEST_ERROR,
			},
		})
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(devResponse.StatusCode)
	w.Write(body)
}
