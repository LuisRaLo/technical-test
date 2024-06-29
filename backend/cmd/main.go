package main

import (
	"context"
	"net/http"
	"technical-challenge/internal/adapters/controllers"
	"technical-challenge/internal/core/application"
	"technical-challenge/internal/core/domain"
	"technical-challenge/swagger"

	repositoriesImpl "technical-challenge/internal/adapters/repositories"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"go.uber.org/zap"
)

// @title Swagger Technical Challenge API
// @version 1.0
// @description REST API for the Technical Challenge
// @termsOfService http://swagger.io/terms/

// @contact.name Luis Enrique Ramírez López
// @contact.url http://www.swagger.io/support
// @contact.email luian.ramirez.12@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	logger, err := repositoriesImpl.NewLogger()
	if err != nil {
		logger.Error("Error creating logger")
	}

	var mux *http.ServeMux = http.NewServeMux()
	var ctx context.Context = context.Background()

	//REPOSITORIES

	//USE CASES
	var sellUseCase domain.SellUseCase = application.NewSellUseCase(logger)

	//CONTROLLERS
	var sellController domain.SellController = controllers.NewSellController(logger, sellUseCase)

	//ROUTES
	setupRoutes(ctx, logger, mux, sellController)

	//SERVER
	logger.Info("Server running on port 8080")
	http.ListenAndServe(":8080", mux)

}

// Configurar rutas
func setupRoutes(ctx context.Context, logger *zap.SugaredLogger, router *http.ServeMux, sellController domain.SellController) {
	//usersEndpointPath := os.Getenv("USERS_ENDPOINT_PATH")

	router.HandleFunc("POST /api/v1/sell", sellController.Sell())
	logger.Info("POST /api/v1/sell endpoint created")

	//swagger
	swagger.SwaggerInfo.Title = "Swagger Technical Challenge API"
	swagger.SwaggerInfo.Description = "This is a sample server Petstore server."
	router.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))

}
