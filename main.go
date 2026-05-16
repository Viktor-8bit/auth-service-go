package main

import (
	"auth-service/endpoints"
	"auth-service/repositories"
	"auth-service/services"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := gin.Default()

	authRepository := repositories.NewAuthRepository("postgres://postgres:postgres@localhost:5432/authservicedb", logger)

	authService := services.NewAuthService(authRepository, logger)
	authEndpoint := endpoints.NewAuthEndpoint(authService, logger)

	r.POST("/login", authEndpoint.Login)
	r.POST("/register", authEndpoint.Register)
	r.Run()
}
