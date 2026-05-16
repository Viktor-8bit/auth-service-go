package main

import (
	"auth-service/endpoints"
	"auth-service/repositories"
	"auth-service/services"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	// load envs
	DBConnect, _ := os.LookupEnv("DBConnect")
	JWTToken, _ := os.LookupEnv("JWTToken")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := gin.Default()

	authRepository := repositories.NewAuthRepository(DBConnect, logger)

	authService := services.NewAuthService(authRepository, logger, JWTToken)
	authEndpoint := endpoints.NewAuthEndpoint(authService, logger)

	r.POST("/login", authEndpoint.Login)
	r.POST("/register", authEndpoint.Register)
	r.Run()
}
