package services

import (
	"auth-service/models"
	"auth-service/repositories"
	"context"
	"log/slog"
)

type RefreshService struct {
	RefreshRepository *repositories.RefreshRepository
	logger            *slog.Logger
}

func NewRefreshService(refreshRepo *repositories.RefreshRepository, logger *slog.Logger) *RefreshService {

	return &RefreshService{
		logger:            logger,
		RefreshRepository: refreshRepo,
	}
}

func (re *RefreshService) LogRefreshToeken(refToken *models.RefreshToken, c context.Context) (error, *int) {

	err, id := re.RefreshRepository.LogNewRefreshToken(refToken, c)

	if err != nil {
		return err, nil
	}

	return nil, id
}
