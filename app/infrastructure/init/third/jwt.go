package third

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}
