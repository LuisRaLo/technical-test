package utils

import (
	"encoding/json"
	"net/http"
	"technical-challenge/internal/core/domain/constants"
	"technical-challenge/internal/core/domain/models"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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

func ValidateBody(w http.ResponseWriter, r *http.Request, payload interface{}, logger *zap.SugaredLogger) []string {
	var errors []string = make([]string, 0)
	validate := validator.New()

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		errors = append(errors, "Invalid JSON")
	}

	err = validate.Struct(payload)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationErr := range validationErrors {
			switch validationErr.Tag() {
			case "required":
				errors = append(errors, validationErr.Field()+" is required")
			case "min":
				errors = append(errors, validationErr.Field()+" must be greater than "+validationErr.Param())
			case "max":
				errors = append(errors, validationErr.Field()+" must be less than "+validationErr.Param())
			case "alphanum":
				errors = append(errors, validationErr.Field()+" must be alphanumeric")
			}
		}
	}

	return errors
}
