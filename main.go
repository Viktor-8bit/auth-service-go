package main

import (
	"auth-service/endpoints"
	"auth-service/repositories"
	"auth-service/services"
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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
	JWTRefresh, _ := os.LookupEnv("JWTRefresh")
	JWTAccessToken, _ := os.LookupEnv("JWTAccessToken")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := gin.Default()

	// db connect
	dbpool, err := pgxpool.New(context.Background(), DBConnect)

	if err != nil {
		logger.Error("Ошибка подключения", "error", err)
	} else {
		logger.Info("Успешное подключние")
	}

	authRepository := repositories.NewAuthRepository(dbpool, logger)
	refreshRepository := repositories.NewRefreshRepository(dbpool, logger)

	refreshService := services.NewRefreshService(refreshRepository, logger)
	authService := services.NewAuthService(authRepository, refreshService, logger, JWTRefresh, JWTAccessToken)
	authEndpoint := endpoints.NewAuthEndpoint(authService, logger)

	r.POST("/login", authEndpoint.Login)
	r.POST("/register", authEndpoint.Register)
	r.POST("/getaccesstoken", authEndpoint.GetAccessToken)
	r.Run()
}
