package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtType string

const (
	JwtRefreshType JwtType = "refresh"
	JwtAccessType  JwtType = "access"
)

type JWTRefreshRequest struct {
	JwtRefresh string `json:"jwtrefresh"`
}

type RefreshClaim struct {
	jwt.RegisteredClaims
	Login string
	Type  JwtType
}

type AccessClaim struct {
	jwt.RegisteredClaims
	Login  string
	Type   JwtType
	RoleId int
}
