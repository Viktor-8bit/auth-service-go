package models

type ErrInvalidCredentials struct {
	Message string
}

func (e ErrInvalidCredentials) Error() string {
	return e.Message
}

func NewErrInvalidCredentials() error {
	return ErrInvalidCredentials{Message: "invalid credentials"}
}

type ErrUserNotFound struct {
	Message string
}

func (e ErrUserNotFound) Error() string {
	return e.Message
}

func NewErrUserNotFound() error {
	return ErrUserNotFound{Message: "user not found"}
}

type ErrUserAlreadyExists struct {
	Message string
}

func (e ErrUserAlreadyExists) Error() string {
	return e.Message
}

func NewErrUserAlreadyExists() error {
	return ErrUserAlreadyExists{Message: "user already exists"}
}

type ErrInvalidEmailFormat struct {
	Message string
}

func (e ErrInvalidEmailFormat) Error() string {
	return e.Message
}

func NewErrInvalidEmailFormat() error {
	return ErrInvalidEmailFormat{Message: "invalid email format"}
}

type ErrPasswordTooShort struct {
	Message string
}

func (e ErrPasswordTooShort) Error() string {
	return e.Message
}

func NewErrPasswordTooShort() error {
	return ErrPasswordTooShort{Message: "password is too short"}
}

type ErrTokenExpired struct {
	Message string
}

func (e ErrTokenExpired) Error() string {
	return e.Message
}

func NewErrTokenExpired() error {
	return ErrTokenExpired{Message: "token has expired"}
}

type ErrInvalidToken struct {
	Message string
}

func (e ErrInvalidToken) Error() string {
	return e.Message
}

func NewErrInvalidToken() error {
	return ErrInvalidToken{Message: "invalid token"}
}

type ErrTokenMissing struct {
	Message string
}

func (e ErrTokenMissing) Error() string {
	return e.Message
}

func NewErrTokenMissing() error {
	return ErrTokenMissing{Message: "token is missing"}
}

type ErrDatabaseConnection struct {
	Message string
}

func (e ErrDatabaseConnection) Error() string {
	return e.Message
}

func NewErrDatabaseConnection() error {
	return ErrDatabaseConnection{Message: "failed to connect to database"}
}

type ErrFailedToCreateUser struct {
	Message string
}

func (e ErrFailedToCreateUser) Error() string {
	return e.Message
}

func NewErrFailedToCreateUser() error {
	return ErrFailedToCreateUser{Message: "failed to create user"}
}

type ErrFailedToRetrieveUser struct {
	Message string
}

func (e ErrFailedToRetrieveUser) Error() string {
	return e.Message
}

func NewErrFailedToRetrieveUser() error {
	return ErrFailedToRetrieveUser{Message: "failed to retrieve user"}
}

type ErrInvalidRefreshToken struct {
	Message string
}

func (e ErrInvalidRefreshToken) Error() string {
	return e.Message
}

func NewErrInvalidRefreshToken() error {
	return ErrInvalidRefreshToken{Message: "invalid refresh token"}
}

type ErrUnauthorizedAccess struct {
	Message string
}

func (e ErrUnauthorizedAccess) Error() string {
	return e.Message
}

func NewErrUnauthorizedAccess() error {
	return ErrUnauthorizedAccess{Message: "unauthorized access"}
}
