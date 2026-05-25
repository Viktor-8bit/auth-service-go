package repositories

import (
	"auth-service/models"
	"context"
	"errors"
	"log/slog"

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

func (ar *AuthRepository) RegisterUser(user *models.User, ctx context.Context) error {

	_, err := ar.DbPool.Exec(ctx, "INSERT INTO users (user_name, passwd_hash, role, mail, salt) VALUES ($1, $2, $3, $4, $5);", user.UserName, user.PasswordHash, user.Role, user.Mail, user.Salt)

	if err != nil {
		ar.logger.Error(err.Error())
		return errors.New("Ошибка")
	}

	return nil
}

func (ar *AuthRepository) GetUserByLogin(login string, ctx context.Context) (*models.User, error) {

	rows, err := ar.DbPool.Query(ctx, "select * from users where user_name=$1", login)

	if err != nil {
		ar.logger.Error(err.Error())
		return nil, errors.New("Ошибка")
	}

	searchedLogin, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("Пользователь не найден")
		}

		ar.logger.Error(err.Error())
		return nil, errors.New("Ошибка")
	}

	return &searchedLogin, nil

}
