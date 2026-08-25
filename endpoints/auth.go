package endpoints

import (
	"auth-service/models"
	"auth-service/services"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthEndpoint struct {
	authService *services.AuthService
	logger      *slog.Logger
}

func NewAuthEndpoint(authService *services.AuthService, Logger *slog.Logger) *AuthEndpoint {
	return &AuthEndpoint{
		authService: authService,
		logger:      Logger,
	}
}

func (ae *AuthEndpoint) Login(c *gin.Context) {

	var req *models.LoginRequest

	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильно заполнены поля формы"})
		return
	}

	token, err := ae.authService.LoginUser(req, ctx)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"refreshtoken": token})
	}
}

func (ae *AuthEndpoint) Register(c *gin.Context) {
	var req *models.RegisterRequest

	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильно заполнены поля формы"})
		return
	}

	// Проверяем наличие пользователя
	user, err := ae.authService.GetUserByLogin(req.Login, ctx)

	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь уже есть в базе"})
		return
	}

	err = ae.authService.RegisterUser(req, ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusAccepted, gin.H{"status": "Ok"})
	}

}

func (ae *AuthEndpoint) GetAccessToken(c *gin.Context) {

	ctx := c.Request.Context()

	var jwtrefresh models.JWTRefreshRequest

	if err := c.ShouldBindJSON(&jwtrefresh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильно заполнены поля формы"})
		return
	}

	resp, err := ae.authService.GetAccessToken(jwtrefresh.JwtRefresh, ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"accesstoken": resp})
}
