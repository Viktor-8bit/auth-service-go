package services

import (
	"auth-service/models"
	"auth-service/repositories"
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

// ______________________ HASH Passwd part ______________________
const (
	saltLength = 16       // length of salt in bytes
	keyLength  = 32       // length of derived key (e.g., for AES-256)
	argontime  = 3        // number of iterations
	memory     = 64 << 10 // memory cost in KiB (~64MB)
	threads    = 4        // number of parallel threads
)

func verifyPassword(password string, salt []byte, expectedHash []byte) bool {
	newHash := argon2.IDKey([]byte(password), salt, argontime, memory, threads, keyLength)

	return subtle.ConstantTimeCompare(newHash, expectedHash) == 1
}

func hashPassword(password string) ([]byte, []byte, error) {
	salt, err := generateSalt(saltLength)
	if err != nil {
		return nil, nil, err
	}

	hashed := argon2.IDKey([]byte(password), salt, argontime, memory, threads, keyLength)

	return hashed, salt, nil
}

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// ______________________ end HASH Passwd part ______________________

type AuthService struct {
	AuthRepository *repositories.AuthRepository
	logger         *slog.Logger
	jwtsecreet     []byte
}

func NewAuthService(authRepo *repositories.AuthRepository, Logger *slog.Logger, secreet string) *AuthService {
	return &AuthService{
		AuthRepository: authRepo,
		logger:         Logger,
		jwtsecreet:     []byte(secreet),
	}
}

func (au *AuthService) RegisterUser(user *models.RegisterRequest, c *gin.Context) error {

	if user.Password != user.PasswordConf {
		return errors.New("The passwords do not match")
	}

	hash, salt, _ := hashPassword(user.Password)

	regUser := models.NewUser(&user.Login, &user.Email, salt, hash)

	err := au.AuthRepository.RegisterUser(regUser, c)

	if err != nil {
		return err
	}

	return nil
}

func (au *AuthService) GetUserByLogin(login string, c *gin.Context) (*models.User, error) {

	user, err := au.AuthRepository.GetUserByLogin(login, c)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (au *AuthService) LoginUser(loginUser *models.LoginRequest, c *gin.Context) (*string, error) {

	// технически уже проверяет наличие пользователя в системе
	user, err := au.GetUserByLogin(loginUser.Login, c)

	if err != nil {
		return nil, err
	}

	res := verifyPassword(loginUser.Password, user.Salt, user.PasswordHash)

	if res == true {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"login": user.UserName,
			"mail":  user.Mail,
			"role":  user.Role,
			"nbf":   time.Now().Add(24 * time.Hour).Unix(),
		})
		tokenString, err := token.SignedString(au.jwtsecreet)

		if err != nil {
			return nil, err
		}

		return &tokenString, nil
	}

	return nil, errors.New("неверный пароль")
}
