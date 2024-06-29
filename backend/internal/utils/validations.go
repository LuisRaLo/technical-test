package utils

import (
	"encoding/json"
	"net/http"
	"unicode"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

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

func ValidateEmail(email string) string {

	validate := validator.New()
	err := validate.Var(email, "email")
	if err != nil {
		return "Invalid email"
	}

	return ""
}

// verify if password is equal to retyped password
func ValidatePassword(password string, retypedPassword string) string {
	if password != retypedPassword {
		return "Passwords do not match"
	}

	// Verificar longitud mínima
	if len(password) < 6 {
		return "Password must be at least 6 characters long"
	}

	//verificar que no tenga espacios ni caracteres especiales sólo letras, números, _, -, ., @, !, ?, #
	for _, char := range password {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '_' && char != '-' && char != '.' && char != '@' && char != '!' && char != '?' && char != '#' {
			return "Password must contain only letters, numbers, _, -, ., @, !, ?, #"
		}
	}

	return ""
}
