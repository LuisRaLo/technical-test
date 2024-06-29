package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"technical-challenge/internal/adapters/controllers"
	"technical-challenge/internal/core/application"
	"technical-challenge/internal/core/datasources"
	"technical-challenge/internal/core/domain"
	"technical-challenge/internal/core/domain/repositories"
	repositoriesImpl "technical-challenge/internal/core/repositories"
	"technical-challenge/internal/middlewares"
	"technical-challenge/internal/utils"

	"technical-challenge/swagger"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/joho/godotenv"
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

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file => ", err)
	}

	logger, err := repositoriesImpl.NewLogger()
	if err != nil {
		logger.Error("Error creating logger")
	}

	var mux *http.ServeMux = http.NewServeMux()
	var ctx context.Context = context.Background()

	//DATASOURCES
	var database1Connection *sql.DB = datasources.Database1Connection(logger)
	firebaseConnection, err := utils.GetFirebaseSession(ctx, logger)
	if err != nil {
		logger.Error("Error getting firebase session")
	}

	//REPOSITORIES
	usersRepository := repositoriesImpl.NewUsersRepository(logger, database1Connection)
	bondRepository := repositoriesImpl.NewBondRepository(logger, database1Connection)

	//USE CASES
	var sellUseCase domain.SellUseCase = application.NewSellUseCase(logger, bondRepository)
	var usersUseCase repositories.UserUseCase = application.NewUsersUseCase(logger, usersRepository, firebaseConnection)

	//CONTROLLERS
	var sellController domain.SellController = controllers.NewSellController(logger, sellUseCase)
	var usersController repositories.UserController = controllers.NewUsersController(logger, usersUseCase)

	//ROUTES
	setupRoutes(ctx, logger, mux, sellController, usersController)

	//SERVER
	logger.Info("Server running on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		logger.Error("Error starting server")
	}

	defer datasources.CloseDB(database1Connection, logger)
	defer logger.Sync()

}

// Configurar rutas
func setupRoutes(
	ctx context.Context,
	logger *zap.SugaredLogger,
	router *http.ServeMux,
	sellController domain.SellController,
	usersController repositories.UserController,
) {
	//usersEndpointPath := os.Getenv("USERS_ENDPOINT_PATH")

	//MIDDLEWARES
	var authorizerMiddleware *middlewares.AuthorizerMiddleware = middlewares.NewAuthorizerMiddleware(logger)

	//swagger
	swagger.SwaggerInfo.Schemes = []string{"http", "https"}
	router.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))

	router.HandleFunc("POST /api/v1/users", usersController.SingUp(ctx))
	logger.Info("POST /api/v1/users endpoint created")

	router.HandleFunc("GET /api/v1/users/byJWT", func(w http.ResponseWriter, r *http.Request) {
		authorizerMiddleware.Authorizer(usersController.GetUserByJWT()).ServeHTTP(w, r)
	})
	logger.Info("GET /api/v1/users/byJWT endpoint created")

	router.HandleFunc("GET /api/v1/users/{user_id}", func(w http.ResponseWriter, r *http.Request) {
		authorizerMiddleware.Authorizer(usersController.GetUserByID(ctx)).ServeHTTP(w, r)
	})
	logger.Info("GET /api/v1/users/{user_id} endpoint created")

	router.HandleFunc("POST /api/v1/sell", func(w http.ResponseWriter, r *http.Request) {
		authorizerMiddleware.Authorizer(sellController.Sell()).ServeHTTP(w, r)
	})
	logger.Info("POST /api/v1/sell endpoint created")
}
