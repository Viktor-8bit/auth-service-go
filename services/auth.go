package services

import (
	"auth-service/models"
	"auth-service/repositories"
	"context"
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"log/slog"
	"time"

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
	jwtrefresh     []byte
	jwtaccess      []byte
}

func NewAuthService(authRepo *repositories.AuthRepository, Logger *slog.Logger, jwtrefresh string, jwtaccess string) *AuthService {
	return &AuthService{
		AuthRepository: authRepo,
		logger:         Logger,
		jwtrefresh:     []byte(jwtrefresh),
		jwtaccess:      []byte(jwtaccess),
	}
}

func (au *AuthService) RegisterUser(user *models.RegisterRequest, c context.Context) error {

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

func (au *AuthService) GetUserByLogin(login string, c context.Context) (*models.User, error) {

	user, err := au.AuthRepository.GetUserByLogin(login, c)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// отправит refresh токен на месяц
func (au *AuthService) LoginUser(loginUser *models.LoginRequest, c context.Context) (*string, error) {

	// технически уже проверяет наличие пользователя в системе

	user, err := au.GetUserByLogin(loginUser.Login, c)

	if err != nil {
		return nil, err
	}

	res := verifyPassword(loginUser.Password, user.Salt, user.PasswordHash)

	claim := models.RefreshClaim{
		Login: *user.UserName,
		Type:  models.JwtRefreshType,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 1, 0)),
		},
	}

	if res == true {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

		tokenString, err := token.SignedString(au.jwtrefresh)

		if err != nil {
			return nil, err
		}

		return &tokenString, nil
	}

	return nil, errors.New("неверный пароль")
}

func (au *AuthService) GetAccessToken(jwtrefresh string, c context.Context) (*string, error) {

	claims := &models.RefreshClaim{}

	token, err := jwt.ParseWithClaims(
		jwtrefresh,
		claims,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(au.jwtrefresh), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid || claims.Type != models.JwtRefreshType {
		return nil, errors.New("invalid token")
	}

	if claims.RegisteredClaims.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, errors.New("token expiret")
	}

	user, err := au.GetUserByLogin(claims.Login, c)

	if err != nil {
		return nil, errors.New("problem with user")
	}

	claim := models.AccessClaim{
		Login:  claims.Login,
		Type:   models.JwtAccessType,
		RoleId: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}

	refreshtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	refreshtokenString, err := refreshtoken.SignedString(au.jwtaccess)

	return &refreshtokenString, nil
}
