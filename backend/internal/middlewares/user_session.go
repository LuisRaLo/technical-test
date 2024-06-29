package middlewares

import (
	"context"
	"net/http"
	"technical-challenge/internal/core/domain/constants"
	"technical-challenge/internal/core/domain/models"
	"technical-challenge/internal/utils"

	"go.uber.org/zap"
)

type AuthorizerMiddleware struct {
	logger *zap.SugaredLogger
}

func NewAuthorizerMiddleware(logger *zap.SugaredLogger) *AuthorizerMiddleware {
	return &AuthorizerMiddleware{
		logger: logger,
	}
}
func (a *AuthorizerMiddleware) Authorizer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userData, err := utils.IsValidSession(ctx, a.logger, r)
		if err != nil {
			res := models.DevResponse{
				StatusCode: http.StatusUnauthorized,
				Response: models.Response401WithResult{
					Message: constants.REQUEST_UNAUTHORIZED,
					Details: []string{err.Error()},
				},
			}
			utils.Response(w, res)
			return
		}

		a.logger.Info("Req /" + r.Method + " " + r.URL.Path + " from user: " + userData.TokenData.UID)
		ctx = context.WithValue(ctx, constants.AuthenticatedUserKey, userData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Function to get userData from the context
func GetUserDataFromContext(ctx context.Context) utils.ResultFirebase {
	if userData, ok := ctx.Value(constants.AuthenticatedUserKey).(utils.ResultFirebase); ok {
		return userData
	}
	return utils.ResultFirebase{}
}
