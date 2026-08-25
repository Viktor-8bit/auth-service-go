package repositories

import (
	"auth-service/models"
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshRepository struct {
	DbPool *pgxpool.Pool
	logger *slog.Logger
}

func NewRefreshRepository(DbPool *pgxpool.Pool, Logger *slog.Logger) *RefreshRepository {

	return &RefreshRepository{
		DbPool: DbPool,
		logger: Logger,
	}
}

func (re *RefreshRepository) GetLogByTokenId(id *int, c context.Context) (error, *models.RefreshToken) {

	rows, err := re.DbPool.Query(c, "select * from refreshtokens where jti=$1;", id)

	if err != nil {
		re.logger.Error(err.Error())
		return errors.New("Ошибка"), nil
	}

	searchedRefreshToken, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.RefreshToken])

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("Токен не найден"), nil
		}

		re.logger.Error(err.Error())
		return errors.New("Ошибка"), nil
	}

	return nil, &searchedRefreshToken
}

func (re *RefreshRepository) LogNewRefreshToken(refToken *models.RefreshToken, c context.Context) (error, *int) {

	var jti *int

	err := re.DbPool.QueryRow(
		c,
		`INSERT INTO refreshtokens
			(user_id, iat, exp, revoked_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING jti`,
		refToken.UserID,
		refToken.IAT,
		refToken.EXP,
		refToken.RevokedAt,
		refToken.CreatedAt,
	).Scan(&jti)

	if err != nil {
		re.logger.Error(err.Error())
		return errors.New("Ошибка"), nil
	}

	return nil, jti
}

//    Column   |           Type           | Collation | Nullable |                  Default
// ------------+--------------------------+-----------+----------+--------------------------------------------
//  user_id    | integer                  |           | not null |
//  iat        | timestamp with time zone |           | not null |
//  exp        | timestamp with time zone |           | not null |
//  revoked_at | timestamp with time zone |           |          |
//  created_at | timestamp with time zone |           | not null | now()
//  jti        | bigint                   |           | not null | nextval('refreshtokens_jti_seq'::regclass)
