package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

/*
*

  - IsValidSession

  - @param ateServices api.Services

  - @param r *http.Request

  - @return models.ResultFirebase

  - This function validates the session of the user

  - by invoking the lambda function LAMBDA_SECURITY_CORE

  - and checking the Authorization header of the request

  - The function returns a models.ResultFirebase struct

  - with the user data if the session is valid

  - If the session is not valid, the function returns an

  - empty models.ResultFirebase struct

  - If there is an error invoking the lambda function

  - or unmarshalling the response, the function returns

  - an empty models.ResultFirebase struct
*/
func IsValidSession(context context.Context, logger *zap.SugaredLogger, r *http.Request) (ResultFirebase, error) {
	var userData ResultFirebase = ResultFirebase{}
	autorization := r.Header.Get("Authorization")
	if autorization == "" {
		logger.Info("Authorization header is empty")
		return userData, fmt.Errorf("authorization header is empty")

	}
	var token string

	//buscar si el token tienen el prefijo Bearer
	if !strings.HasPrefix(autorization, "Bearer") {
		token = autorization
	}

	token = strings.TrimPrefix(autorization, "Bearer ")

	authData, err := GetUserDataByToken(context, logger, token)
	if err != nil {
		logger.Infow("Error getting user data by token: ", err)
	}
	return authData, err

}
