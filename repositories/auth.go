package repositories

import (
	"auth-service/models"
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	DbPool *pgxpool.Pool
	logger *slog.Logger
}

func NewAuthRepository(connStr string, Logger *slog.Logger) *AuthRepository {

	dbpool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		Logger.Error("Ошибка подключения", "error", err)
	} else {
		Logger.Info("Успешное подключние")
	}

	// defer dbpool.Close()

	return &AuthRepository{
		DbPool: dbpool,
		logger: Logger,
	}
}

func (ar *AuthRepository) RegisterUser(user *models.User, c *gin.Context) error {

	_, err := ar.DbPool.Query(c, "INSERT INTO users (user_name, passwd_hash, role, mail, salt) VALUES ($1, $2, $3, $4, $5);", user.UserName, user.PasswordHash, user.Role, user.Mail, user.Salt)

	if err != nil {
		return err
	}

	return nil
}

func (ar *AuthRepository) GetUserByLogin(login string, c *gin.Context) (*models.User, error) {

	rows, err := ar.DbPool.Query(c, "select * from users where user_name=$1", login)

	if err != nil {
		return nil, err
	}

	searchedLogin, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {
		return nil, err
	}

	return &searchedLogin, nil

}
